package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	DB, err := sql.Open("sqlite3", "api.db")
	if err != nil {
		panic("could not connect to database")
	}

	DB.SetMaxOpenConns(10) // controll how many connections can be open
	DB.SetMaxIdleConns(5) // how many connections to keep open when none are being used
}