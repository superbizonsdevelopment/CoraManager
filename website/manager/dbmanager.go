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

	_, err = Base.Exec("CREATE TABLE IF NOT EXISTS users (id Integer NOT NULL PRIMARY KEY AUTOINCREMENT, name varchar(255), password varchar(255), email varchar(255), permissions varchar(255), orders varchar(255));")

	if err != nil {
		return err
	}

	_, err = Base.Exec("CREATE TABLE IF NOT EXISTS orders (id Integer NOT NULL PRIMARY KEY AUTOINCREMENT, name varchar(255), content varchar(255), purchaser integer, createdatdate varchar(255), status varchar(255));")

	if err != nil {
		return err
	}

	Base = database
	return nil
}
