package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Import the PostgreSQL driver package for Go.

	"github.com/Imlucky883/simple_bank/api"
	db "github.com/Imlucky883/simple_bank/db/sqlc"
)

const serverAddress = ":8080"

func main() {
	conn, err := sql.Open("postgres", "postgresql://postgres:postgres@localhost:5432/simple_bank_test?sslmode=disable")
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store) // Replace 'nil' with the appropriate argument if needed

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
