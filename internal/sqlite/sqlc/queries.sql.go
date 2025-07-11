// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.29.0
// source: queries.sql

package sqlc

import (
	"context"
	"database/sql"
	"time"
)

const addTeamMember = `-- name: AddTeamMember :exec

INSERT INTO team_members (team_id, user_id, role)
VALUES (?, ?, ?)
`

type AddTeamMemberParams struct {
	TeamID int64  `json:"team_id"`
	UserID int64  `json:"user_id"`
	Role   string `json:"role"`
}

// Team Members
// Add a member to a team
func (q *Queries) AddTeamMember(ctx context.Context, arg AddTeamMemberParams) error {
	_, err := q.exec(ctx, q.addTeamMemberStmt, addTeamMember, arg.TeamID, arg.UserID, arg.Role)
	return err
}

const addTeamSource = `-- name: AddTeamSource :exec

INSERT INTO team_sources (team_id, source_id)
VALUES (?, ?)
`

type AddTeamSourceParams struct {
	TeamID   int64 `json:"team_id"`
	SourceID int64 `json:"source_id"`
}

// Team Sources
// Add a data source to a team
func (q *Queries) AddTeamSource(ctx context.Context, arg AddTeamSourceParams) error {
	_, err := q.exec(ctx, q.addTeamSourceStmt, addTeamSource, arg.TeamID, arg.SourceID)
	return err
}

const countAdminUsers = `-- name: CountAdminUsers :one
SELECT COUNT(*) FROM users WHERE role = ? AND status = ?
`

type CountAdminUsersParams struct {
	Role   string `json:"role"`
	Status string `json:"status"`
}

// Count active admin users
func (q *Queries) CountAdminUsers(ctx context.Context, arg CountAdminUsersParams) (int64, error) {
	row := q.queryRow(ctx, q.countAdminUsersStmt, countAdminUsers, arg.Role, arg.Status)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const countUserSessions = `-- name: CountUserSessions :one
SELECT COUNT(*) FROM sessions WHERE user_id = ? AND expires_at > ?
`

type CountUserSessionsParams struct {
	UserID    int64     `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
}

// Count active sessions for a user
func (q *Queries) CountUserSessions(ctx context.Context, arg CountUserSessionsParams) (int64, error) {
	row := q.queryRow(ctx, q.countUserSessionsStmt, countUserSessions, arg.UserID, arg.ExpiresAt)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const createAPIToken = `-- name: CreateAPIToken :one

INSERT INTO api_tokens (user_id, name, token_hash, prefix, expires_at)
VALUES (?, ?, ?, ?, ?)
RETURNING id
`

type CreateAPITokenParams struct {
	UserID    int64        `json:"user_id"`
	Name      string       `json:"name"`
	TokenHash string       `json:"token_hash"`
	Prefix    string       `json:"prefix"`
	ExpiresAt sql.NullTime `json:"expires_at"`
}

// API Tokens
// Create a new API token
func (q *Queries) CreateAPIToken(ctx context.Context, arg CreateAPITokenParams) (int64, error) {
	row := q.queryRow(ctx, q.createAPITokenStmt, createAPIToken,
		arg.UserID,
		arg.Name,
		arg.TokenHash,
		arg.Prefix,
		arg.ExpiresAt,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createSession = `-- name: CreateSession :exec

INSERT INTO sessions (id, user_id, expires_at, created_at)
VALUES (?, ?, ?, ?)
`

type CreateSessionParams struct {
	ID        string    `json:"id"`
	UserID    int64     `json:"user_id"`
	ExpiresAt time.Time `json:"expires_at"`
	CreatedAt time.Time `json:"created_at"`
}

// Sessions
// Create a new session
func (q *Queries) CreateSession(ctx context.Context, arg CreateSessionParams) error {
	_, err := q.exec(ctx, q.createSessionStmt, createSession,
		arg.ID,
		arg.UserID,
		arg.ExpiresAt,
		arg.CreatedAt,
	)
	return err
}

const createSource = `-- name: CreateSource :one

INSERT INTO sources (
    name, _meta_is_auto_created, _meta_ts_field, _meta_severity_field, host, username, password, database, table_name, description, ttl_days, created_at, updated_at
) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, datetime('now'), datetime('now'))
RETURNING id
`

type CreateSourceParams struct {
	Name              string         `json:"name"`
	MetaIsAutoCreated int64          `json:"_meta_is_auto_created"`
	MetaTsField       string         `json:"_meta_ts_field"`
	MetaSeverityField sql.NullString `json:"_meta_severity_field"`
	Host              string         `json:"host"`
	Username          string         `json:"username"`
	Password          string         `json:"password"`
	Database          string         `json:"database"`
	TableName         string         `json:"table_name"`
	Description       sql.NullString `json:"description"`
	TtlDays           int64          `json:"ttl_days"`
}

// Sources
// Create a new source entry
func (q *Queries) CreateSource(ctx context.Context, arg CreateSourceParams) (int64, error) {
	row := q.queryRow(ctx, q.createSourceStmt, createSource,
		arg.Name,
		arg.MetaIsAutoCreated,
		arg.MetaTsField,
		arg.MetaSeverityField,
		arg.Host,
		arg.Username,
		arg.Password,
		arg.Database,
		arg.TableName,
		arg.Description,
		arg.TtlDays,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createTeam = `-- name: CreateTeam :one

INSERT INTO teams (name, description)
VALUES (?, ?)
RETURNING id
`

type CreateTeamParams struct {
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
}

// Teams
// Create a new team
func (q *Queries) CreateTeam(ctx context.Context, arg CreateTeamParams) (int64, error) {
	row := q.queryRow(ctx, q.createTeamStmt, createTeam, arg.Name, arg.Description)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createTeamSourceQuery = `-- name: CreateTeamSourceQuery :one

INSERT INTO team_queries (team_id, source_id, name, description, query_type, query_content)
VALUES (?, ?, ?, ?, ?, ?)
RETURNING id
`

type CreateTeamSourceQueryParams struct {
	TeamID       int64          `json:"team_id"`
	SourceID     int64          `json:"source_id"`
	Name         string         `json:"name"`
	Description  sql.NullString `json:"description"`
	QueryType    string         `json:"query_type"`
	QueryContent string         `json:"query_content"`
}

// Team Queries
// Create a new query for a team and source
func (q *Queries) CreateTeamSourceQuery(ctx context.Context, arg CreateTeamSourceQueryParams) (int64, error) {
	row := q.queryRow(ctx, q.createTeamSourceQueryStmt, createTeamSourceQuery,
		arg.TeamID,
		arg.SourceID,
		arg.Name,
		arg.Description,
		arg.QueryType,
		arg.QueryContent,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const createUser = `-- name: CreateUser :one

INSERT INTO users (email, full_name, role, status, last_login_at)
VALUES (?, ?, ?, ?, ?)
RETURNING id
`

type CreateUserParams struct {
	Email       string       `json:"email"`
	FullName    string       `json:"full_name"`
	Role        string       `json:"role"`
	Status      string       `json:"status"`
	LastLoginAt sql.NullTime `json:"last_login_at"`
}

// Users
// Create a new user
func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (int64, error) {
	row := q.queryRow(ctx, q.createUserStmt, createUser,
		arg.Email,
		arg.FullName,
		arg.Role,
		arg.Status,
		arg.LastLoginAt,
	)
	var id int64
	err := row.Scan(&id)
	return id, err
}

const deleteAPIToken = `-- name: DeleteAPIToken :exec
DELETE FROM api_tokens WHERE id = ? AND user_id = ?
`

type DeleteAPITokenParams struct {
	ID     int64 `json:"id"`
	UserID int64 `json:"user_id"`
}

// Delete an API token by ID and user ID (ensure user owns the token)
func (q *Queries) DeleteAPIToken(ctx context.Context, arg DeleteAPITokenParams) error {
	_, err := q.exec(ctx, q.deleteAPITokenStmt, deleteAPIToken, arg.ID, arg.UserID)
	return err
}

const deleteExpiredAPITokens = `-- name: DeleteExpiredAPITokens :exec
DELETE FROM api_tokens WHERE expires_at IS NOT NULL AND expires_at < datetime('now')
`

// Delete all expired API tokens
func (q *Queries) DeleteExpiredAPITokens(ctx context.Context) error {
	_, err := q.exec(ctx, q.deleteExpiredAPITokensStmt, deleteExpiredAPITokens)
	return err
}

const deleteSession = `-- name: DeleteSession :exec
DELETE FROM sessions WHERE id = ?
`

// Delete a session by ID
func (q *Queries) DeleteSession(ctx context.Context, id string) error {
	_, err := q.exec(ctx, q.deleteSessionStmt, deleteSession, id)
	return err
}

const deleteSource = `-- name: DeleteSource :exec
DELETE FROM sources WHERE id = ?
`

// Delete a source by ID
func (q *Queries) DeleteSource(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteSourceStmt, deleteSource, id)
	return err
}

const deleteTeam = `-- name: DeleteTeam :exec
DELETE FROM teams WHERE id = ?
`

// Delete a team by ID
func (q *Queries) DeleteTeam(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteTeamStmt, deleteTeam, id)
	return err
}

const deleteTeamSourceQuery = `-- name: DeleteTeamSourceQuery :exec
DELETE FROM team_queries
WHERE id = ? AND team_id = ? AND source_id = ?
`

type DeleteTeamSourceQueryParams struct {
	ID       int64 `json:"id"`
	TeamID   int64 `json:"team_id"`
	SourceID int64 `json:"source_id"`
}

// Delete a query by ID for a specific team and source
func (q *Queries) DeleteTeamSourceQuery(ctx context.Context, arg DeleteTeamSourceQueryParams) error {
	_, err := q.exec(ctx, q.deleteTeamSourceQueryStmt, deleteTeamSourceQuery, arg.ID, arg.TeamID, arg.SourceID)
	return err
}

const deleteUser = `-- name: DeleteUser :exec
DELETE FROM users WHERE id = ?
`

// Delete a user by ID
func (q *Queries) DeleteUser(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.deleteUserStmt, deleteUser, id)
	return err
}

const deleteUserSessions = `-- name: DeleteUserSessions :exec
DELETE FROM sessions WHERE user_id = ?
`

// Delete all sessions for a user
func (q *Queries) DeleteUserSessions(ctx context.Context, userID int64) error {
	_, err := q.exec(ctx, q.deleteUserSessionsStmt, deleteUserSessions, userID)
	return err
}

const getAPIToken = `-- name: GetAPIToken :one
SELECT id, user_id, name, token_hash, prefix, last_used_at, expires_at, created_at, updated_at FROM api_tokens WHERE id = ?
`

// Get an API token by ID
func (q *Queries) GetAPIToken(ctx context.Context, id int64) (ApiToken, error) {
	row := q.queryRow(ctx, q.getAPITokenStmt, getAPIToken, id)
	var i ApiToken
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.TokenHash,
		&i.Prefix,
		&i.LastUsedAt,
		&i.ExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getAPITokenByHash = `-- name: GetAPITokenByHash :one
SELECT id, user_id, name, token_hash, prefix, last_used_at, expires_at, created_at, updated_at FROM api_tokens WHERE token_hash = ?
`

// Get an API token by its hash (for authentication)
func (q *Queries) GetAPITokenByHash(ctx context.Context, tokenHash string) (ApiToken, error) {
	row := q.queryRow(ctx, q.getAPITokenByHashStmt, getAPITokenByHash, tokenHash)
	var i ApiToken
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.Name,
		&i.TokenHash,
		&i.Prefix,
		&i.LastUsedAt,
		&i.ExpiresAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSession = `-- name: GetSession :one
SELECT id, user_id, expires_at, created_at FROM sessions WHERE id = ?
`

// Get a session by ID
func (q *Queries) GetSession(ctx context.Context, id string) (Session, error) {
	row := q.queryRow(ctx, q.getSessionStmt, getSession, id)
	var i Session
	err := row.Scan(
		&i.ID,
		&i.UserID,
		&i.ExpiresAt,
		&i.CreatedAt,
	)
	return i, err
}

const getSource = `-- name: GetSource :one
SELECT id, name, _meta_is_auto_created, _meta_ts_field, _meta_severity_field, host, username, password, "database", table_name, description, ttl_days, created_at, updated_at FROM sources WHERE id = ?
`

// Get a single source by ID
func (q *Queries) GetSource(ctx context.Context, id int64) (Source, error) {
	row := q.queryRow(ctx, q.getSourceStmt, getSource, id)
	var i Source
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.MetaIsAutoCreated,
		&i.MetaTsField,
		&i.MetaSeverityField,
		&i.Host,
		&i.Username,
		&i.Password,
		&i.Database,
		&i.TableName,
		&i.Description,
		&i.TtlDays,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getSourceByName = `-- name: GetSourceByName :one
SELECT id, name, _meta_is_auto_created, _meta_ts_field, _meta_severity_field, host, username, password, "database", table_name, description, ttl_days, created_at, updated_at FROM sources WHERE database = ? AND table_name = ?
`

type GetSourceByNameParams struct {
	Database  string `json:"database"`
	TableName string `json:"table_name"`
}

// Get a single source by table name and database
func (q *Queries) GetSourceByName(ctx context.Context, arg GetSourceByNameParams) (Source, error) {
	row := q.queryRow(ctx, q.getSourceByNameStmt, getSourceByName, arg.Database, arg.TableName)
	var i Source
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.MetaIsAutoCreated,
		&i.MetaTsField,
		&i.MetaSeverityField,
		&i.Host,
		&i.Username,
		&i.Password,
		&i.Database,
		&i.TableName,
		&i.Description,
		&i.TtlDays,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTeam = `-- name: GetTeam :one
SELECT id, name, description, created_at, updated_at FROM teams WHERE id = ?
`

// Get a team by ID
func (q *Queries) GetTeam(ctx context.Context, id int64) (Team, error) {
	row := q.queryRow(ctx, q.getTeamStmt, getTeam, id)
	var i Team
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTeamByName = `-- name: GetTeamByName :one
SELECT id, name, description, created_at, updated_at FROM teams WHERE name = ?
`

// Get a team by its name
func (q *Queries) GetTeamByName(ctx context.Context, name string) (Team, error) {
	row := q.queryRow(ctx, q.getTeamByNameStmt, getTeamByName, name)
	var i Team
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Description,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getTeamMember = `-- name: GetTeamMember :one
SELECT team_id, user_id, role, created_at FROM team_members WHERE team_id = ? AND user_id = ?
`

type GetTeamMemberParams struct {
	TeamID int64 `json:"team_id"`
	UserID int64 `json:"user_id"`
}

// Get a team member
func (q *Queries) GetTeamMember(ctx context.Context, arg GetTeamMemberParams) (TeamMember, error) {
	row := q.queryRow(ctx, q.getTeamMemberStmt, getTeamMember, arg.TeamID, arg.UserID)
	var i TeamMember
	err := row.Scan(
		&i.TeamID,
		&i.UserID,
		&i.Role,
		&i.CreatedAt,
	)
	return i, err
}

const getTeamSourceQuery = `-- name: GetTeamSourceQuery :one
SELECT id, team_id, source_id, name, description, query_type, query_content, created_at, updated_at FROM team_queries
WHERE id = ? AND team_id = ? AND source_id = ?
`

type GetTeamSourceQueryParams struct {
	ID       int64 `json:"id"`
	TeamID   int64 `json:"team_id"`
	SourceID int64 `json:"source_id"`
}

// Get a query by ID for a specific team and source
func (q *Queries) GetTeamSourceQuery(ctx context.Context, arg GetTeamSourceQueryParams) (TeamQuery, error) {
	row := q.queryRow(ctx, q.getTeamSourceQueryStmt, getTeamSourceQuery, arg.ID, arg.TeamID, arg.SourceID)
	var i TeamQuery
	err := row.Scan(
		&i.ID,
		&i.TeamID,
		&i.SourceID,
		&i.Name,
		&i.Description,
		&i.QueryType,
		&i.QueryContent,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, email, full_name, role, status, last_login_at, last_active_at, created_at, updated_at FROM users WHERE id = ?
`

// Get a user by ID
func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.queryRow(ctx, q.getUserStmt, getUser, id)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FullName,
		&i.Role,
		&i.Status,
		&i.LastLoginAt,
		&i.LastActiveAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const getUserByEmail = `-- name: GetUserByEmail :one
SELECT id, email, full_name, role, status, last_login_at, last_active_at, created_at, updated_at FROM users WHERE email = ?
`

// Get a user by email
func (q *Queries) GetUserByEmail(ctx context.Context, email string) (User, error) {
	row := q.queryRow(ctx, q.getUserByEmailStmt, getUserByEmail, email)
	var i User
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.FullName,
		&i.Role,
		&i.Status,
		&i.LastLoginAt,
		&i.LastActiveAt,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listAPITokensForUser = `-- name: ListAPITokensForUser :many
SELECT id, user_id, name, token_hash, prefix, last_used_at, expires_at, created_at, updated_at FROM api_tokens WHERE user_id = ? ORDER BY created_at DESC
`

// List all API tokens for a user
func (q *Queries) ListAPITokensForUser(ctx context.Context, userID int64) ([]ApiToken, error) {
	rows, err := q.query(ctx, q.listAPITokensForUserStmt, listAPITokensForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ApiToken{}
	for rows.Next() {
		var i ApiToken
		if err := rows.Scan(
			&i.ID,
			&i.UserID,
			&i.Name,
			&i.TokenHash,
			&i.Prefix,
			&i.LastUsedAt,
			&i.ExpiresAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listQueriesByTeamAndSource = `-- name: ListQueriesByTeamAndSource :many
SELECT id, team_id, source_id, name, description, query_type, query_content, created_at, updated_at FROM team_queries WHERE team_id = ? AND source_id = ? ORDER BY created_at DESC
`

type ListQueriesByTeamAndSourceParams struct {
	TeamID   int64 `json:"team_id"`
	SourceID int64 `json:"source_id"`
}

// List all queries for a specific team and source
func (q *Queries) ListQueriesByTeamAndSource(ctx context.Context, arg ListQueriesByTeamAndSourceParams) ([]TeamQuery, error) {
	rows, err := q.query(ctx, q.listQueriesByTeamAndSourceStmt, listQueriesByTeamAndSource, arg.TeamID, arg.SourceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TeamQuery{}
	for rows.Next() {
		var i TeamQuery
		if err := rows.Scan(
			&i.ID,
			&i.TeamID,
			&i.SourceID,
			&i.Name,
			&i.Description,
			&i.QueryType,
			&i.QueryContent,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSourceTeams = `-- name: ListSourceTeams :many
SELECT t.id, t.name, t.description, t.created_at, t.updated_at
FROM teams t
JOIN team_sources ts ON t.id = ts.team_id
WHERE ts.source_id = ?
ORDER BY t.name
`

// List all teams a data source is a member of
func (q *Queries) ListSourceTeams(ctx context.Context, sourceID int64) ([]Team, error) {
	rows, err := q.query(ctx, q.listSourceTeamsStmt, listSourceTeams, sourceID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Team{}
	for rows.Next() {
		var i Team
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSources = `-- name: ListSources :many
SELECT id, name, _meta_is_auto_created, _meta_ts_field, _meta_severity_field, host, username, password, "database", table_name, description, ttl_days, created_at, updated_at FROM sources ORDER BY created_at DESC
`

// Get all sources ordered by creation date
func (q *Queries) ListSources(ctx context.Context) ([]Source, error) {
	rows, err := q.query(ctx, q.listSourcesStmt, listSources)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Source{}
	for rows.Next() {
		var i Source
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.MetaIsAutoCreated,
			&i.MetaTsField,
			&i.MetaSeverityField,
			&i.Host,
			&i.Username,
			&i.Password,
			&i.Database,
			&i.TableName,
			&i.Description,
			&i.TtlDays,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listSourcesForUser = `-- name: ListSourcesForUser :many
SELECT DISTINCT s.id, s.name, s._meta_is_auto_created, s._meta_ts_field, s._meta_severity_field, s.host, s.username, s.password, s."database", s.table_name, s.description, s.ttl_days, s.created_at, s.updated_at FROM sources s
JOIN team_sources ts ON s.id = ts.source_id
JOIN team_members tm ON ts.team_id = tm.team_id
WHERE tm.user_id = ?
ORDER BY s.created_at DESC
`

// List all sources a user has access to
func (q *Queries) ListSourcesForUser(ctx context.Context, userID int64) ([]Source, error) {
	rows, err := q.query(ctx, q.listSourcesForUserStmt, listSourcesForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Source{}
	for rows.Next() {
		var i Source
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.MetaIsAutoCreated,
			&i.MetaTsField,
			&i.MetaSeverityField,
			&i.Host,
			&i.Username,
			&i.Password,
			&i.Database,
			&i.TableName,
			&i.Description,
			&i.TtlDays,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTeamMembers = `-- name: ListTeamMembers :many
SELECT tm.team_id, tm.user_id, tm.role, tm.created_at
FROM team_members tm
WHERE tm.team_id = ?
ORDER BY tm.created_at
`

// List all members of a team
func (q *Queries) ListTeamMembers(ctx context.Context, teamID int64) ([]TeamMember, error) {
	rows, err := q.query(ctx, q.listTeamMembersStmt, listTeamMembers, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []TeamMember{}
	for rows.Next() {
		var i TeamMember
		if err := rows.Scan(
			&i.TeamID,
			&i.UserID,
			&i.Role,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTeamMembersWithDetails = `-- name: ListTeamMembersWithDetails :many
SELECT tm.team_id, tm.user_id, tm.role, tm.created_at, u.email, u.full_name
FROM team_members tm
JOIN users u ON tm.user_id = u.id
WHERE tm.team_id = ?
ORDER BY tm.created_at ASC
`

type ListTeamMembersWithDetailsRow struct {
	TeamID    int64     `json:"team_id"`
	UserID    int64     `json:"user_id"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
	Email     string    `json:"email"`
	FullName  string    `json:"full_name"`
}

// List all members of a team with user details
func (q *Queries) ListTeamMembersWithDetails(ctx context.Context, teamID int64) ([]ListTeamMembersWithDetailsRow, error) {
	rows, err := q.query(ctx, q.listTeamMembersWithDetailsStmt, listTeamMembersWithDetails, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListTeamMembersWithDetailsRow{}
	for rows.Next() {
		var i ListTeamMembersWithDetailsRow
		if err := rows.Scan(
			&i.TeamID,
			&i.UserID,
			&i.Role,
			&i.CreatedAt,
			&i.Email,
			&i.FullName,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTeamSources = `-- name: ListTeamSources :many
SELECT s.id, s.name, s._meta_is_auto_created, s._meta_ts_field, s._meta_severity_field, s.host, s.username, s.password, s."database", s.table_name, s.description, s.ttl_days, s.created_at, s.updated_at
FROM sources s
JOIN team_sources ts ON s.id = ts.source_id
WHERE ts.team_id = ?
ORDER BY s.created_at DESC
`

// List all data sources in a team
func (q *Queries) ListTeamSources(ctx context.Context, teamID int64) ([]Source, error) {
	rows, err := q.query(ctx, q.listTeamSourcesStmt, listTeamSources, teamID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Source{}
	for rows.Next() {
		var i Source
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.MetaIsAutoCreated,
			&i.MetaTsField,
			&i.MetaSeverityField,
			&i.Host,
			&i.Username,
			&i.Password,
			&i.Database,
			&i.TableName,
			&i.Description,
			&i.TtlDays,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTeams = `-- name: ListTeams :many
SELECT t.id, t.name, t.description, t.created_at, t.updated_at, COUNT(tm.user_id) as member_count
FROM teams t
LEFT JOIN team_members tm ON t.id = tm.team_id
GROUP BY t.id
ORDER BY t.created_at DESC
`

type ListTeamsRow struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	MemberCount int64          `json:"member_count"`
}

// List all teams
func (q *Queries) ListTeams(ctx context.Context) ([]ListTeamsRow, error) {
	rows, err := q.query(ctx, q.listTeamsStmt, listTeams)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListTeamsRow{}
	for rows.Next() {
		var i ListTeamsRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.MemberCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listTeamsForUser = `-- name: ListTeamsForUser :many
SELECT
    t.id,
    t.name,
    t.description,
    t.created_at,
    t.updated_at,
    tm.role,  -- The current user's role in this team
    (SELECT COUNT(*) FROM team_members sub_tm WHERE sub_tm.team_id = t.id) as member_count
FROM
    teams t
JOIN
    team_members tm ON t.id = tm.team_id
WHERE
    tm.user_id = ?  -- The current user ID
ORDER BY
    t.created_at DESC
`

type ListTeamsForUserRow struct {
	ID          int64          `json:"id"`
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	Role        string         `json:"role"`
	MemberCount int64          `json:"member_count"`
}

// List all teams a user is a member of
func (q *Queries) ListTeamsForUser(ctx context.Context, userID int64) ([]ListTeamsForUserRow, error) {
	rows, err := q.query(ctx, q.listTeamsForUserStmt, listTeamsForUser, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []ListTeamsForUserRow{}
	for rows.Next() {
		var i ListTeamsForUserRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
			&i.Role,
			&i.MemberCount,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUserTeams = `-- name: ListUserTeams :many
SELECT t.id, t.name, t.description, t.created_at, t.updated_at
FROM teams t
JOIN team_members tm ON t.id = tm.team_id
WHERE tm.user_id = ?
ORDER BY t.name
`

// List all teams a user is a member of
func (q *Queries) ListUserTeams(ctx context.Context, userID int64) ([]Team, error) {
	rows, err := q.query(ctx, q.listUserTeamsStmt, listUserTeams, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Team{}
	for rows.Next() {
		var i Team
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const listUsers = `-- name: ListUsers :many
SELECT id, email, full_name, role, status, last_login_at, last_active_at, created_at, updated_at FROM users ORDER BY created_at ASC
`

// List all users
func (q *Queries) ListUsers(ctx context.Context) ([]User, error) {
	rows, err := q.query(ctx, q.listUsersStmt, listUsers)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []User{}
	for rows.Next() {
		var i User
		if err := rows.Scan(
			&i.ID,
			&i.Email,
			&i.FullName,
			&i.Role,
			&i.Status,
			&i.LastLoginAt,
			&i.LastActiveAt,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const removeTeamMember = `-- name: RemoveTeamMember :exec
DELETE FROM team_members
WHERE team_id = ? AND user_id = ?
`

type RemoveTeamMemberParams struct {
	TeamID int64 `json:"team_id"`
	UserID int64 `json:"user_id"`
}

// Remove a member from a team
func (q *Queries) RemoveTeamMember(ctx context.Context, arg RemoveTeamMemberParams) error {
	_, err := q.exec(ctx, q.removeTeamMemberStmt, removeTeamMember, arg.TeamID, arg.UserID)
	return err
}

const removeTeamSource = `-- name: RemoveTeamSource :exec
DELETE FROM team_sources WHERE team_id = ? AND source_id = ?
`

type RemoveTeamSourceParams struct {
	TeamID   int64 `json:"team_id"`
	SourceID int64 `json:"source_id"`
}

// Remove a data source from a team
func (q *Queries) RemoveTeamSource(ctx context.Context, arg RemoveTeamSourceParams) error {
	_, err := q.exec(ctx, q.removeTeamSourceStmt, removeTeamSource, arg.TeamID, arg.SourceID)
	return err
}

const teamHasSource = `-- name: TeamHasSource :one

SELECT COUNT(*) FROM team_sources
WHERE team_id = ? AND source_id = ?
`

type TeamHasSourceParams struct {
	TeamID   int64 `json:"team_id"`
	SourceID int64 `json:"source_id"`
}

// Additional queries for user-source and team-source access
// Check if a team has access to a source
func (q *Queries) TeamHasSource(ctx context.Context, arg TeamHasSourceParams) (int64, error) {
	row := q.queryRow(ctx, q.teamHasSourceStmt, teamHasSource, arg.TeamID, arg.SourceID)
	var count int64
	err := row.Scan(&count)
	return count, err
}

const updateAPITokenLastUsed = `-- name: UpdateAPITokenLastUsed :exec
UPDATE api_tokens
SET last_used_at = datetime('now'),
    updated_at = datetime('now')
WHERE id = ?
`

// Update the last used timestamp for an API token
func (q *Queries) UpdateAPITokenLastUsed(ctx context.Context, id int64) error {
	_, err := q.exec(ctx, q.updateAPITokenLastUsedStmt, updateAPITokenLastUsed, id)
	return err
}

const updateSource = `-- name: UpdateSource :exec
UPDATE sources
SET name = ?,
    _meta_is_auto_created = ?,
    _meta_ts_field = ?,
    _meta_severity_field = ?,
    host = ?,
    username = ?,
    password = ?,
    database = ?,
    table_name = ?,
    description = ?,
    ttl_days = ?,
    updated_at = datetime('now')
WHERE id = ?
`

type UpdateSourceParams struct {
	Name              string         `json:"name"`
	MetaIsAutoCreated int64          `json:"_meta_is_auto_created"`
	MetaTsField       string         `json:"_meta_ts_field"`
	MetaSeverityField sql.NullString `json:"_meta_severity_field"`
	Host              string         `json:"host"`
	Username          string         `json:"username"`
	Password          string         `json:"password"`
	Database          string         `json:"database"`
	TableName         string         `json:"table_name"`
	Description       sql.NullString `json:"description"`
	TtlDays           int64          `json:"ttl_days"`
	ID                int64          `json:"id"`
}

// Update an existing source
func (q *Queries) UpdateSource(ctx context.Context, arg UpdateSourceParams) error {
	_, err := q.exec(ctx, q.updateSourceStmt, updateSource,
		arg.Name,
		arg.MetaIsAutoCreated,
		arg.MetaTsField,
		arg.MetaSeverityField,
		arg.Host,
		arg.Username,
		arg.Password,
		arg.Database,
		arg.TableName,
		arg.Description,
		arg.TtlDays,
		arg.ID,
	)
	return err
}

const updateTeam = `-- name: UpdateTeam :exec
UPDATE teams
SET name = ?,
    description = ?,
    updated_at = ?
WHERE id = ?
`

type UpdateTeamParams struct {
	Name        string         `json:"name"`
	Description sql.NullString `json:"description"`
	UpdatedAt   time.Time      `json:"updated_at"`
	ID          int64          `json:"id"`
}

// Update a team
func (q *Queries) UpdateTeam(ctx context.Context, arg UpdateTeamParams) error {
	_, err := q.exec(ctx, q.updateTeamStmt, updateTeam,
		arg.Name,
		arg.Description,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}

const updateTeamMemberRole = `-- name: UpdateTeamMemberRole :exec
UPDATE team_members
SET role = ?
WHERE team_id = ? AND user_id = ?
`

type UpdateTeamMemberRoleParams struct {
	Role   string `json:"role"`
	TeamID int64  `json:"team_id"`
	UserID int64  `json:"user_id"`
}

// Update a team member's role
func (q *Queries) UpdateTeamMemberRole(ctx context.Context, arg UpdateTeamMemberRoleParams) error {
	_, err := q.exec(ctx, q.updateTeamMemberRoleStmt, updateTeamMemberRole, arg.Role, arg.TeamID, arg.UserID)
	return err
}

const updateTeamSourceQuery = `-- name: UpdateTeamSourceQuery :exec
UPDATE team_queries
SET name = ?,
    description = ?,
    query_type = ?,
    query_content = ?,
    updated_at = datetime('now')
WHERE id = ? AND team_id = ? AND source_id = ?
`

type UpdateTeamSourceQueryParams struct {
	Name         string         `json:"name"`
	Description  sql.NullString `json:"description"`
	QueryType    string         `json:"query_type"`
	QueryContent string         `json:"query_content"`
	ID           int64          `json:"id"`
	TeamID       int64          `json:"team_id"`
	SourceID     int64          `json:"source_id"`
}

// Update a query for a team and source
func (q *Queries) UpdateTeamSourceQuery(ctx context.Context, arg UpdateTeamSourceQueryParams) error {
	_, err := q.exec(ctx, q.updateTeamSourceQueryStmt, updateTeamSourceQuery,
		arg.Name,
		arg.Description,
		arg.QueryType,
		arg.QueryContent,
		arg.ID,
		arg.TeamID,
		arg.SourceID,
	)
	return err
}

const updateUser = `-- name: UpdateUser :exec
UPDATE users
SET email = ?,
    full_name = ?,
    role = ?,
    status = ?,
    last_login_at = ?,
    last_active_at = ?,
    updated_at = ?
WHERE id = ?
`

type UpdateUserParams struct {
	Email        string       `json:"email"`
	FullName     string       `json:"full_name"`
	Role         string       `json:"role"`
	Status       string       `json:"status"`
	LastLoginAt  sql.NullTime `json:"last_login_at"`
	LastActiveAt sql.NullTime `json:"last_active_at"`
	UpdatedAt    time.Time    `json:"updated_at"`
	ID           int64        `json:"id"`
}

// Update a user
func (q *Queries) UpdateUser(ctx context.Context, arg UpdateUserParams) error {
	_, err := q.exec(ctx, q.updateUserStmt, updateUser,
		arg.Email,
		arg.FullName,
		arg.Role,
		arg.Status,
		arg.LastLoginAt,
		arg.LastActiveAt,
		arg.UpdatedAt,
		arg.ID,
	)
	return err
}

const userHasSourceAccess = `-- name: UserHasSourceAccess :one
SELECT COUNT(*) FROM team_members tm
JOIN team_sources ts ON tm.team_id = ts.team_id
WHERE tm.user_id = ? AND ts.source_id = ?
`

type UserHasSourceAccessParams struct {
	UserID   int64 `json:"user_id"`
	SourceID int64 `json:"source_id"`
}

// Check if a user has access to a source through any team
func (q *Queries) UserHasSourceAccess(ctx context.Context, arg UserHasSourceAccessParams) (int64, error) {
	row := q.queryRow(ctx, q.userHasSourceAccessStmt, userHasSourceAccess, arg.UserID, arg.SourceID)
	var count int64
	err := row.Scan(&count)
	return count, err
}
