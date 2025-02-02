package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error
	testDB, err = sql.Open("postgres", "postgresql://postgres:postgres@localhost:5432/testdb?sslmode=disable")
	if err != nil {
		log.Fatal("Cannot connect to database:", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
