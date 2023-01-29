package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"github.com/ruffiano/simplebank/api"
	db "github.com/ruffiano/simplebank/db/sqlc"
	"github.com/ruffiano/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server: ", err)
	}
}
