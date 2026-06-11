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

mock:
	mockgen -source=db/sqlc/store.go -destination=db/mock/store.go -package=mockdb -aux_files=github.com/DefinitelyNotJay/flyte/db/sqlc=db/sqlc/querier.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc server test mock