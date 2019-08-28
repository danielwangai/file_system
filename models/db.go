package models

import (
	"database/sql"
	"log"

	// enable usage of postgres
	_ "github.com/lib/pq"
)

type Database struct {
	DB *sql.DB
}

// Connect establishes connection to the database
func (database *Database) Connect(host, port, user, password, dbName, sslMode string) error {
	connString := "host=" + host + " port=" + port + " user=" + user + " password=" + password + " dbname=" + dbName + " sslmode=" + sslMode
	var err error
	database.DB, err = sql.Open("postgres", connString)
	if err != nil {
		log.Fatal("error connecting to the database", err)
		return err
	}
	return nil
}
