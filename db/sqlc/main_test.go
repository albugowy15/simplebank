package db

import (
	"context"
	"github.com/albugowy15/simplebank/utils"
	"log"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
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
