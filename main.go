package main

import (
	"database/sql"
	"log"

	"gita_app/api"
	db "gita_app/db/sqlc"

	_ "github.com/lib/pq"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://suresh:2003@localhost:5432/gita_db?sslmode=disable"
	serverAddress = "0.0.0.0:8081"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatalf("cannot connect to db: %v", err)
	}

	store := db.New(conn)
	server := api.NewServer(store)

	log.Printf("starting server at %s", serverAddress)
	err = server.Start(serverAddress)
	if err != nil {
		log.Fatalf("cannot start server: %v", err)
	}
}
