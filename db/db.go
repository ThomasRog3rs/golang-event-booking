package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDb() {
	DB, err := sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("Could not connect to databse.")
	}

	DB.SetMaxOpenConns(10)
}
