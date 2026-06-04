postgres:
	docker run --name postgres13 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres13-alpine

createdb:
	docker exec -it postgres13 createdb --username=root --owner=root flyte

dropdb:
	docker exec -it postgres13 dropdb flyte

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5433/flyte?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5433/flyte?sslmode=disable" -verbose down
sqlc:
	sqlc generate
server:
	go run main.go

test:
	go test -v -cover ./...