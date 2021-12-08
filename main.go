package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/Threx-code/simple_bank/api"
	db "github.com/Threx-code/simple_bank/db/sqlc"
	"github.com/Threx-code/simple_bank/utils"
)

// const (
// 	dbDriver      = "postgres"
// 	dbSource      = "postgresql://root:secret@localhost:5654/simple_bank?sslmode=disable"
// 	serverAddress = "0.0.0.0:8020"
// )

func main() {

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server, err := api.NewServer(config, store)
	if err != nil {
		log.Fatal("cannot create server", err)
	}

	err = server.Start(config.ServerAddress)

	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
