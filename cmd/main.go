package main

import (
	"database/sql"
	"go-rest-api/cmd/api"
	"go-rest-api/config"
	"go-rest-api/db"
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	db, err := db.NewMysqlStorage(&mysql.Config{
		User:                 config.EnvVars.DBUser,
		Passwd:               config.EnvVars.DBPassword,
		Addr:                 config.EnvVars.DBAddress,
		DBName:               config.EnvVars.DBName,
		AllowNativePasswords: true,
		ParseTime:            true,
	})
	if err != nil {
		log.Fatal(err)
	}

	initStorage(db)

	server := api.NewAPIServer(":9090", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}

func initStorage(db *sql.DB) {
	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	log.Println("db: connected")
}
