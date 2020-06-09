package controllers

import (
	"FoodsRefactored/models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//Controller is
type Controller struct{}

var foods []models.Food

func (c Controller) GetFoods(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var food models.Food
		//assign an empty slice every time method is called
		foods = []models.Food{}

		rows, err := database.Query("select * from items")
		checkErr(err)
		defer rows.Close()

		for rows.Next() {
			rows.Scan(&food.ID, &food.Title, &food.URL, &food.Likes, &food.UserID)
			checkErr(err)

			foods = append(foods, food)
		}

		json.NewEncoder(w).Encode(foods)
	}
}

//what just happened here?
//remove getFoods name to get an anonymous funciton
//then add it to the above method to return??
func (c Controller) GetFood(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var food Food

		food := models.Food{}
		params := mux.Vars(r)

		rows := database.QueryRow("select * from items where id=?;", params["id"])

		err := rows.Scan(&food.ID, &food.Title, &food.URL, &food.Likes, &food.UserID)
		checkErr(err)
		// handle if no entry is found
		json.NewEncoder(w).Encode(food)
	}
}

func (c Controller) AddFood(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//create variable type food
		var food models.Food

		//fetching the parameter r
		//decoding the food
		json.NewDecoder(r.Body).Decode(&food)

		//makig variable for querying into the database sqlite
		query := "insert into items(title, url, likes, userid) values (?,?,?,?)"

		//check for the error(database, statement, etc, whatever it is), and must close it afterwards
		statement, err := database.Prepare(query)
		checkErr(err)
		defer statement.Close()

		result, err := statement.Exec(&food.Title, &food.URL, &food.Likes, &food.UserID)

		id, err := result.LastInsertId()
		checkErr(err)

		json.NewEncoder(w).Encode(id)
	}
}

func (c Controller) UpdateFood(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var food models.Food
		json.NewDecoder(r.Body).Decode(&food)

		query := "update items set title = ?, url = ?, likes = ?, userid = ? where id = ?"

		statement, err := database.Prepare(query)
		checkErr(err)
		defer statement.Close()

		result, err := statement.Exec(&food.Title, &food.URL, &food.Likes, &food.UserID, &food.ID)

		rowsUpdated, err := result.RowsAffected()

		//id, err := result.LastInsertId()
		checkErr(err)

		json.NewEncoder(w).Encode(rowsUpdated)
	}
}

func (c Controller) RemoveFood(database *sql.DB) http.HandlerFunc { //food := Food{}
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

func checkErr(err error) {
	if err != nil {
		log.Println("Yup it returned an error", err)
		//log.Fatal(err)
	}
}
