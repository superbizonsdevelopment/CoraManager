package manager

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var (
	Base *sql.DB
)

func ConnectToDatabase() error {
	database, err := sql.Open("sqlite3", "./website.db")

	if err != nil {
		return err
	}

	Base = database
}
