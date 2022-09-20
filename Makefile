postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root myshop-db
dropdb:
	docker exec -it postgres12 dropdb myshop-db

migrationfie:
	migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/myshop-db?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://postgres:password@localhost:5432/myshop-db?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go

.PHONY: createdb dropdb migrateup migratedown sqlc test server
