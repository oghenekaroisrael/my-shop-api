package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/oghenekaroisrael/myshopapi/api"
	db "github.com/oghenekaroisrael/myshopapi/db/sqlc"
	"github.com/oghenekaroisrael/myshopapi/utils"
)

func main() {
	utils.LoadConfig()
	config := utils.GetEnvWithKey
	DBDriver := config("DB_DRIVER")
	DBSource := config("DB_SOURCE")
	ServerAddress := config("SERVER_ADDRESS")

	conn, err := sql.Open(DBDriver, DBSource)
	if err != nil {
		log.Fatal("cannot connect to test db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(*store)
	err = server.Start(ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
