package driver

import (
	"FoodsRefactored/utils"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() *sql.DB {

	database, err := sql.Open("sqlite3", "/Users/kritisharma/projects/Databases/foods.db")
	utils.CheckErr(err)


	statement, err := database.Prepare("create table if not exists items (id integer primary key, title text, url text, likes integer, userid integer)")
	utils.CheckErr(err)
	defer statement.Close()

	statement.Exec()

	return database
}
