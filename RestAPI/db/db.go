package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db")

	if err != nil {
		panic("could not connect to databse")
	}
	DB.SetMaxOpenConns(10) // Maximum number of connection
	DB.SetMaxIdleConns(5)  // How many connection we want to keep open if no one is connected (5 in this case)

	createTables()
}

func createTables() {
	createEventsTable := `
		CREATE TABLE IF NOT EXISTS events (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			description TEXT NOT NULL,
			location TEXT NOT NULL,
			datetime DATETIME NOT NULL,
			user_id INTEGER
		)
	`
	_, err := DB.Exec(createEventsTable) // execute the query
	if err != nil {
		panic("could not create table: " + err.Error()) // provide more detailed error information
	}
}
