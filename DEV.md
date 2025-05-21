# Dev

## Pre-requisites

```bash
# Install pnpm
curl -fsSL https://get.pnpm.io/install.sh | sh -

# install just
wget https://github.com/casey/just/releases/download/1.40.0/just-1.40.0-x86_64-unknown-linux-musl.tar.gz
tar -xzf just-1.40.0-x86_64-unknown-linux-musl.tar.gz
sudo mv just /usr/local/bin
```

## Start dev

```bash
just run-backend
just run-frontend
```