package driver

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var database *sql.DB

func ConnectDB() *sql.DB {
	return database
}
