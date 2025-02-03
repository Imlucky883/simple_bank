package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Imlucky883/simple_bank/db/util"
	_ "github.com/lib/pq"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	// Load configuration from app.env
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("❌ Cannot load config: %v", err)
	}

	// Connect to test database
	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatalf("❌ Cannot connect to test database: %v", err)
	}
	defer testDB.Close()

	testQueries = New(testDB)

	// Run tests
	os.Exit(m.Run())
}
