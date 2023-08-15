create_migrate:
	migrate create -ext sql -dir db/migration -seq init_schema

create_postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

start_postgres:
	docker start postgres12

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v "H:\Visual Studio Projects\simplebank:/db" -w /db kjconroy/sqlc generate

test:
	go test -v -cover ./...

.PHONY: create_migrate create_postgres start_postgres createdb dropdb migrateup migratedown sqlc test