package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	var err error
	// No password, just username@tcp(...)
	dsn := "root:@tcp(localhost:3306)/solusphere"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatal("Database ping failed:", err)
	}

	log.Println("Database connected successfully")
}
