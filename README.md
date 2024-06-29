```markdown
# Server Game

Server Game is a project that provides a RESTful API for managing games. It is written in Go and uses the Gin framework.

## Quick Start

- Install [Go](https://golang.org/dl/)
- Install [Docker](https://docs.docker.com/get-docker/)
- SET environment variables in .env file
- SET environment Go

## Clone project

```bash
git clone https://github.com/MinhSang97/Server-Game.git
```

## Move to Server-Game folder

```bash
cd Server-Game
```
## Install container docker

```bash
docker compose up -d
```

## Install dependencies

```bash
go mod tidy
```

## DB Migration

```bash
make migrate-up
```

## Run project

```bash
go run ./cmd/main.go
```

## LICENSE

This project is distributed under the MIT License. See the `LICENSE` file for more information.
```
