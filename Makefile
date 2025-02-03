-- Active: 1738420667066@@127.0.0.1@5432@simple_bank_test

all : postgres createdb migrateup server

postgres:
	docker run -d --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres postgres:17.2-alpine3.21

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres --host=localhost simple_bank_test

dropdb:
	docker exec -it postgres dropdb simple_bank_test

migrateup:
	migrate -path db/migration -database "postgres://postgres:postgres@localhost:5432/simple_bank_test?sslmode=disable" up

migratedown:
	migrate -path db/migration -database "postgres://postgres:postgres@localhost:5432/simple_bank_test?sslmode=disable" down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server
