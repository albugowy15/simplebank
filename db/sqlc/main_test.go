package db

import (
	"database/sql"
	"log"
	"os"
	"simplebank/utils"
	"testing"

	_ "github.com/lib/pq"
)

var (
	testQueries *Queries
	testDB      *sql.DB
)

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("Can't connect to the database:", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
