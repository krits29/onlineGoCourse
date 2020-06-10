package controllers

import (
	"FoodsRefactored/models"
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//Controller is
//type Controller struct{}

var users []models.User

func (c Controller) GetUsers(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		//assign an empty slice every time method is called
		users = []models.User{}

		rows, err := database.Query("select * from users")
		checkErr(err)
		defer rows.Close()

		for rows.Next() {
			rows.Scan(&user.ID, &user.Username, &user.Password)
			checkErr(err)

			users = append(users, user)
		}

		json.NewEncoder(w).Encode(users)
	}
}

func (c Controller) GetUser(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user := models.User{}
		params := mux.Vars(r)

		rows := database.QueryRow("select * from user where id=?;", params["id"])

		err := rows.Scan(&user.ID, &user.Username, &user.Password)
		checkErr(err)
		// handle if no entry is found
		json.NewEncoder(w).Encode(user)
	}
}

func (c Controller) AddUser(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//create variable type food
		var user models.User

		//fetching the parameter r
		//decoding the food
		json.NewDecoder(r.Body).Decode(&user)

		//makig variable for querying into the database sqlite
		query := "insert into users(username, password) values (?,?)"

		//check for the error(database, statement, etc, whatever it is), and must close it afterwards
		statement, err := database.Prepare(query)
		checkErr(err)
		defer statement.Close()

		result, err := statement.Exec(&user.Username, &user.Password)

		id, err := result.LastInsertId()
		checkErr(err)

		json.NewEncoder(w).Encode(id)
	}
}

func (c Controller) UpdateUser(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		json.NewDecoder(r.Body).Decode(&user)

		query := "update users set username = ?, password = ? where id = ?"

		statement, err := database.Prepare(query)
		checkErr(err)
		defer statement.Close()

		result, err := statement.Exec(&user.Username, &user.Password, &user.ID)

		rowsUpdated, err := result.RowsAffected()

		//id, err := result.LastInsertId()
		checkErr(err)

		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

func (c Controller) RemoveUser(database *sql.DB) http.HandlerFunc { //food := Food{}
	return func(w http.ResponseWriter, r *http.Request) {
		//invoing vars on mux to request the object which will return a map with key value pairs
		params := mux.Vars(r)

		statement, err := database.Prepare("delete from items where id=?;")
		checkErr(err)
		defer statement.Close()

		result, err := statement.Exec(params["id"])
		checkErr(err)

		rowsDeleted, err := result.RowsAffected()
		checkErr(err)

		json.NewEncoder(w).Encode(rowsDeleted)
	}
}

// func checkErr(err error) {
// 	if err != nil {
// 		log.Println("Yup it returned an error", err)
// 		//log.Fatal(err)
// 	}
// }
