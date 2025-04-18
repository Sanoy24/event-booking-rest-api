package databse

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func InitializeDb() {
	var err error
	DB, err = sql.Open("sqlite3", "database.db")

	if err != nil {
		panic("can't connect to the database")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTable()
}

func createTable() {
	createEventsTable := `CREATE TABLE IF NOT EXISTS events(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		date_time DATETIME NOT NULL,
		user_id INTEGER	
	)
	`
	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("could not create events table")
	}
}
