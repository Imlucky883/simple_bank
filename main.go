package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq" // Import the PostgreSQL driver package for Go.

	"github.com/Imlucky883/simple_bank/api"
	db "github.com/Imlucky883/simple_bank/db/sqlc"
	"github.com/Imlucky883/simple_bank/db/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store) // Replace 'nil' with the appropriate argument if needed

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Error starting server:", err)
	}
}
