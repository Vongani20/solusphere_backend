package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect(user, password, host, dbname string) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, password, host, dbname)
	database, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	// Test connection
	if err := database.Ping(); err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	DB = database
	log.Println("Connected to MySQL database!")
}
