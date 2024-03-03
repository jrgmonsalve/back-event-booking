package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "./event_booking.db")
	if err != nil {
		panic(err)
	}
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	createTables()
}

func createTables() {
	sqlStmt := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL, 
		description TEXT NOT NULL,
		location TEXT NOT NULL, 
		datetime DATETIME NOT NULL, 
		userid INTEGER
		);
	`
	_, err := DB.Exec(sqlStmt)
	if err != nil {
		panic(err)
	}
}
