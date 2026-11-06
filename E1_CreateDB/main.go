// Creates a PostgreSQL database from Go if it does not already exist.

package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"regexp"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {

	// Connection settings
	host := "localhost"
	port := 5432
	user := "postgres"
	password := ""
	// Name of the database to create
	var databaseName string
	fmt.Print("Enter database name: ")
	fmt.Scan(&databaseName)

	// Allow only letters, digits and underscores to prevent SQL injection
	if !regexp.MustCompile(`^[A-Za-z_][A-Za-z0-9_]*$`).MatchString(databaseName) {
		log.Fatal("invalid database name")
	}

	// Connect to the default "postgres" maintenance database
	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(user, password),
		Host:     fmt.Sprintf("%s:%d", host, port),
		Path:     "/postgres",
		RawQuery: "sslmode=disable",
	}
	adminConn := u.String()

	db, err := sql.Open("pgx", adminConn)
	if err != nil {
		log.Fatalf("failed to create connection: %v", err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		log.Fatalf("could not connect to the server: %v", err)
	}

	// Check whether the database already exists
	var exists bool
	err = db.QueryRow(
		"SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = $1)", databaseName,
	).Scan(&exists)
	if err != nil {
		log.Fatalf("failed to check for existing database: %v", err)
	}

	if exists {
		fmt.Printf("Database %s already exists.\n", databaseName)
		return
	}

	// Create the database using the server's default settings
	createQuery := fmt.Sprintf(`CREATE DATABASE "%s"`, databaseName)
	if _, err := db.Exec(createQuery); err != nil {
		log.Fatalf("failed to create database: %v", err)
	}

	fmt.Printf("Database %s created successfully.\n", databaseName)
}
