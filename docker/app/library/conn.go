package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func GetConnection() (*sqlx.DB, error) {
	user := os.Getenv("POSTGRESDB_USER")
	pwd := os.Getenv("POSTGRESDB_PASSWORD")
	name := os.Getenv("POSTGRESDB_NAME")
	host := os.Getenv("POSTGRESDB_ADDRESS")
	port := os.Getenv("POSTGRESDB_PORT")
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", user, pwd, host, port, name)

	db, err := sqlx.Open("postgres", dbURL)

	if err != nil {
		log.Print("Error while accessing database: " + err.Error())
		return nil, err
	}

	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	return db, nil
}

func CloseConnection(conn *sqlx.DB) {
	err := conn.Close()
	if err != nil {
		log.Fatal(InternalServerError)
	}
}
