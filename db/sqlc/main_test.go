package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/Threx-code/simple_bank/utils"
	_ "github.com/lib/pq"
)

// const (
// 	dbServer = "postgres"
// 	dbSource = "postgresql://root:secret@localhost:5654/simple_bank?sslmode=disable"
// )

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	//var err error

	config, err := utils.LoadConfig("../..")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("we cannot connect to database", err)
	}

	testQueries = New(testDB)
	os.Exit(m.Run())
}
