package main

import (
	"context"
	"flyte/api"
	db "flyte/db/sqlc"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	dbSource      = "postgresql://root:password@localhost:5433/flyte?sslmode=disable"
	serverAddress = "0.0.0.0:8000"
)

func main() {
	connPool, err := pgxpool.New(context.Background(), dbSource)

	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(connPool)
	server := api.Newserver(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}
