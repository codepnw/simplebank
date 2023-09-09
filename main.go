package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"github.com/codepnw/simplebank/api"
	db "github.com/codepnw/simplebank/db/sqlc"
	"github.com/codepnw/simplebank/util"
)

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("connect error to db: ", err)
	}
	
	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cant start server: ", err)
	}
}