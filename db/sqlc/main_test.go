package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/oghenekaroisrael/myshopapi/utils"
)

var testQueries *Queries

func TestMain(m *testing.M) {
	config := utils.GetEnvWithKey
	DBDriver := config("DB_DRIVER")
	DBSource := config("DB_SOURCE")

	conn, err := sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Fatal("cannot connect to test db: ", err)
	}

	testQueries = New(conn)
	os.Exit(m.Run())
}
