package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/microsoft/go-mssqldb"
)

var db *sql.DB

func DbConnection() {
	// Read from environment variables

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	server := os.Getenv("DB_SERVER")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	connString := fmt.Sprintf("user=%s;password=%s;server=%s;port=%s;database=%s;Encrypt=disable;TrustServerCertificate=true", user, password, server, port, dbName)

	// server := os.Getenv("DB_SERVER")
	// port := os.Getenv("DB_PORT")
	// dbName := os.Getenv("DB_NAME")

	// connString := fmt.Sprintf("server=%s;port=%s;database=%s;Encrypt=disable;TrustServerCertificate=true", server, port, dbName)


	var err error
	db, err = sql.Open("sqlserver", connString)
	if err != nil {
		log.Fatal("Error opening connection: ", err.Error())
	}

	// Ping to verify the desktop-based connection
	err = db.Ping()
	if err != nil {
		log.Fatal("Connection failed: ", err.Error())
	}

	fmt.Println("Connected successfully using Windows Authentication!")

	enableAutoCreateTables()
}

