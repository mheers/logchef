package app

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"time"

	"github.com/mr-karan/logchef/internal/auth"
	"github.com/mr-karan/logchef/internal/clickhouse"
	"github.com/mr-karan/logchef/internal/config"
	"github.com/mr-karan/logchef/internal/core"
	"github.com/mr-karan/logchef/internal/server"
	"github.com/mr-karan/logchef/internal/sqlite"
	"github.com/mr-karan/logchef/pkg/logger"
)

// App represents the core application context, holding dependencies and configuration.
type App struct {
	Config     *config.Config
	SQLite     *sqlite.DB
	ClickHouse *clickhouse.Manager
	Logger     *slog.Logger
	server     *server.Server
	WebFS      http.FileSystem
	BuildInfo  string
	Version    string
}

// Options contains configuration needed when creating a new App instance.
type Options struct {
	ConfigPath string
	WebFS      http.FileSystem // Web filesystem for serving static files.
	BuildInfo  string
	Version    string
}

// New creates and configures a new App instance.
func New(opts Options) (*App, error) {
	cfg, err := config.Load(opts.ConfigPath)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	app := &App{
		Config:    cfg,
		Logger:    logger.New(cfg.Logging.Level == "debug"),
		WebFS:     opts.WebFS,
		BuildInfo: opts.BuildInfo,
		Version:   opts.Version,
	}

	return app, nil
}

// Initialize sets up application components like database connections,
// the OIDC provider, and the HTTP server.
func (a *App) Initialize(ctx context.Context) error {
	var err error

	// Initialize SQLite database.
	sqliteOpts := sqlite.Options{
		Config: a.Config.SQLite,
		Logger: a.Logger,
	}
	a.SQLite, err = sqlite.New(sqliteOpts)
	if err != nil {
		return fmt.Errorf("failed to initialize sqlite: %w", err)
	}

	// Initialize admin users based on configuration.
	if err := core.InitAdminUsers(ctx, a.SQLite, a.Logger, a.Config.Auth.AdminEmails); err != nil {
		a.Logger.Error("failed to initialize admin users", "error", err)
		return fmt.Errorf("failed to initialize admin users: %w", err)
	}

	// Initialize ClickHouse connection manager.
	a.ClickHouse = clickhouse.NewManager(a.Logger)

	// Initialize OIDC Provider.
	// This is optional; if OIDC is not configured, auth features relying on it might be disabled.
	oidcProvider, err := auth.NewOIDCProvider(&a.Config.OIDC, a.Logger)
	if err != nil {
		if errors.Is(err, auth.ErrOIDCProviderNotConfigured) {
			a.Logger.Warn("OIDC provider not configured, skipping OIDC setup")
			// oidcProvider will be nil; dependent features should handle this.
		} else {
			return fmt.Errorf("failed to initialize OIDC provider: %w", err)
		}
	}

	// Load existing sources from SQLite into the ClickHouse manager
	// to establish connections for querying.
	sources, err := a.SQLite.ListSources(ctx)
	if err != nil {
		return fmt.Errorf("failed to list sources: %w", err)
	}
	for _, source := range sources {
		a.Logger.Info("initializing source connection",
			"source_id", source.ID,
			"table", source.Connection.TableName)
		if err := a.ClickHouse.AddSource(source); err != nil {
			// Log failure but continue initialization.
			// The health check system will attempt to recover these connections.
			a.Logger.Warn("failed to initialize source connection, will attempt recovery via health checks",
				"source_id", source.ID,
				"error", err)
		}
	}

	// Start background health checks for the ClickHouse manager.
	// Use 0 to trigger the default interval defined in the manager.
	a.ClickHouse.StartBackgroundHealthChecks(0)

	// Initialize HTTP server.
	serverOpts := server.ServerOptions{
		Config:       a.Config,
		SQLite:       a.SQLite,
		ClickHouse:   a.ClickHouse,
		OIDCProvider: oidcProvider,
		FS:           a.WebFS,
		Logger:       a.Logger,
		BuildInfo:    a.BuildInfo,
		Version:      a.Version,
	}
	a.server = server.New(serverOpts)

	return nil
}

// Start begins the application's main execution loop (starts the HTTP server).
func (a *App) Start() error {
	if a.server == nil {
		return fmt.Errorf("server not initialized")
	}
	a.Logger.Info("starting server")
	return a.server.Start()
}

// Shutdown gracefully stops all application components with timeouts.
func (a *App) Shutdown(ctx context.Context) error {
	a.Logger.Info("shutting down application")

	// Ensure a shutdown context with timeout exists.
	if ctx == nil {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
	}

	// Create derived contexts with shorter timeouts for each component
	serverCtx, serverCancel := context.WithTimeout(ctx, 5*time.Second)
	defer serverCancel()

	clickhouseCtx, clickhouseCancel := context.WithTimeout(ctx, 8*time.Second)
	defer clickhouseCancel()
	// Shutdown server first to stop accepting new requests.
	if a.server != nil {
		a.Logger.Info("shutting down HTTP server")

		serverDone := make(chan error, 1)
		go func() {
			serverDone <- a.server.Shutdown(serverCtx)
		}()

		// Wait for server shutdown or timeout
		select {
		case err := <-serverDone:
			if err != nil {
				a.Logger.Error("error shutting down server", "error", err)
			} else {
				a.Logger.Info("HTTP server shut down successfully")
			}
		case <-serverCtx.Done():
			a.Logger.Warn("timeout shutting down HTTP server, continuing")
		}
	}

	// Close ClickHouse manager (stops health checks and closes clients).
	if a.ClickHouse != nil {
		a.Logger.Info("shutting down ClickHouse connections")

		clickhouseDone := make(chan error, 1)
		go func() {
			clickhouseDone <- a.ClickHouse.Close()
		}()

		// Wait for ClickHouse shutdown or timeout
		select {
		case err := <-clickhouseDone:
			if err != nil {
				a.Logger.Error("error closing clickhouse manager", "error", err)
			} else {
				a.Logger.Info("ClickHouse connections closed successfully")
			}
		case <-clickhouseCtx.Done():
			a.Logger.Warn("timeout closing ClickHouse connections, continuing")
		}
	}

	// Close database connections.
	if a.SQLite != nil {
		a.Logger.Info("closing SQLite connection")
		// SQLite should close almost instantly, no need for a separate goroutine
		if err := a.SQLite.Close(); err != nil {
			a.Logger.Error("error closing SQLite", "error", err)
		} else {
			a.Logger.Info("SQLite connection closed successfully")
		}
	}

	a.Logger.Info("application shutdown complete")
	return nil
}
