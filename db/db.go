package db

import (
	"database/sql"
	"log"

	"github.com/go-sql-driver/mysql"
)

func NewMySQLStorage(confg mysql.Config) (*sql.DB, error) {
	db, err := sql.Open("mysql", confg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	return db, nil
}
