package main

import (
	"course/api"
	db "course/db/sqlc"
	"course/util"
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func main() {

	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connection, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(connection)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}

}
