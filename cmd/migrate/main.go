package main

import (
	"errors"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	mysqlDriver "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"go-rest-api/config"
	"go-rest-api/db"
	"log"
	"os"
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

	driver, err := mysqlDriver.WithInstance(db, &mysqlDriver.Config{})
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithDatabaseInstance(
		"file://cmd/migrate/migrations", "mysql", driver)
	if err != nil {
		log.Fatal(err)
	}

	cmd := os.Args[len(os.Args)-1]
	if cmd == "up" {
		if err = m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal(err)
		}
	}

	if cmd == "down" {
		if err = m.Down(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
			log.Fatal(err)
		}
	}
}
