package controllers

import (
	"FoodsRefactored/models"
	"FoodsRefactored/utils"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Controller is
//type Controller struct{}

var users []models.User

//GetUsers is
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

//GetUser is
func (c Controller) GetUser(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		user := models.User{}
		params := mux.Vars(r)

		rows := database.QueryRow("select * from users where id=?;", params["id"])

		user.Password = ""

		err := rows.Scan(&user.ID, &user.Username, &user.Password)
		checkErr(err)
		// handle if no entry is found
		
		json.NewEncoder(w).Encode(user)
	}
}

//plain method -- nothing fancy

//GetUserFromDB is
func GetUserFromDB(database *sql.DB, user *models.User) (*models.User, error) {

	dbUser := &models.User{}
	log.Println(user.Username)
	rows := database.QueryRow("select * from users where username=?;", user.Username)

	err := rows.Scan(&dbUser.ID, &dbUser.Username, &dbUser.Password)
	checkErr(err)
	return dbUser, err
}

//Register is
func (c Controller) Register(database *sql.DB) http.HandlerFunc {
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
		hashedPassword, _ := utils.HashPassword(user.Password)

		result, err := statement.Exec(&user.Username, hashedPassword)

		id, err := result.LastInsertId()
		checkErr(err)

		json.NewEncoder(w).Encode(id)

		log.Println("register successful", user)
	}
}

//UpdateUser is
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

//RemoveUser is
func (c Controller) RemoveUser(database *sql.DB) http.HandlerFunc { //food := Food{}
	return func(w http.ResponseWriter, r *http.Request) {
		//invoing vars on mux to request the object which will return a map with key value pairs
		params := mux.Vars(r)

		statement, err := database.Prepare("delete from users where id=?;")
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
