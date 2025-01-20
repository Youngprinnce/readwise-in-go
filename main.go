package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
)

func main() {
	cfg := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "highlights",
		AllowNativePasswords: true,
		ParseTime:            true,
	}

	// Connect to the database
	storage := NewMySQLStorage(cfg)

	// Create the tables
	db, err := storage.Init()
	if err != nil {
		log.Fatal(err)
	}

	// Create the API server
	apiServer := NewAPIServer(":3000", db)
	apiServer.Run()
}
