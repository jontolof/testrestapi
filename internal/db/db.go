package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

// Init establishes a connection to the database and runs a simple migration
func Init() {
	// Get environment variables with default values
	host := os.Getenv("DB_HOST")
	if host == "" {
		host = "localhost"
	}

	port := os.Getenv("DB_PORT")
	if port == "" {
		port = "5432"
	}

	user := os.Getenv("DB_USER")
	if user == "" {
		user = "myuser"
	}

	password := os.Getenv("DB_PASSWORD")
	if password == "" {
		password = "mypassword"
	}

	dbname := os.Getenv("DB_NAME")
	if dbname == "" {
		dbname = "mydb"
	}

	// Build DSN (Data Source Name)
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	var err error
	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Couldn't open database connection: ", err)
	}

	// Test connection
	err = DB.Ping()
	if err != nil {
		log.Fatal("Couldn't connect to database: ", err)
	}

	// Create table if it does not exist
	_, err = DB.Exec(`
		CREATE TABLE IF NOT EXISTS todos (
			id SERIAL PRIMARY KEY,
			item TEXT NOT NULL,
			completed BOOLEAN NOT NULL DEFAULT false
		);
	`)
	if err != nil {
		log.Fatal("Failed migration: ", err)
	}
}
