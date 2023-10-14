create_migrate:
	migrate create -ext sql -dir db/migration -seq init_schema

create_postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

start_postgres:
	docker start postgres12

create_db:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

drop_db:
	docker exec -it postgres12 dropdb simple_bank

migrate_up:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc_win: #for Windows
	docker run --rm -v "$(pwd):/db" -w /db sqlc/sqlc generate

sqlc_mac: #for macOS
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: create_migrate create_postgres start_postgres create_db drop_db migrate_up migrate_down sqlc_win sqlc_mac test server
