# Go Simple Bank

TBA

## Installation
Download and install Docker Desktop\
https://www.docker.com/products/docker-desktop/

Download and install Golang language\
https://go.dev/doc/install/

### macOS

```bash
brew install golang-migrate
```
```bash
brew install sqlc
```

### Windows

```bash
scoop install migrate
```
```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

## Usage
Check the Makefile file.
- for example run ```make create_migrate``` create a database migration.
- for example run ```create_postgres``` create and start a PostgreSQL 12 database in a Docker container.