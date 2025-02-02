postgres:
	docker run -d --name postgres -p 5432:5432 -e POSTGRES_PASSWORD=postgres postgres:17.2-alpine3.21

createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres --host=localhost laxman

dropdb:
	docker exec -it postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgres://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" up

migratedown:
	migrate -path db/migration -database "postgres://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" down

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc
