package db

import (
	"context"
	"log"
	"os"
	"simplebank/utils"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

var testStore Store

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Println("cannot load config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("Can't connect to the database:", err)
	}

	testStore = NewStore(connPool)
	os.Exit(m.Run())
}
