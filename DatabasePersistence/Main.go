package main

import (
	"database/sql"
	"encoding/json"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"
)

//Food is
type Food struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	URL    string `json:"url"`
	Likes  int    `json:"likes"`
	UserID int    `json:"userid"`
}

var foods []Food

var database *sql.DB

func main() {
	var err error
	database, err = sql.Open("sqlite3", "/Users/kritisharma/projects/Databases/foods.db")
	checkErr(err)
	defer database.Close()

	statement, err := database.Prepare("create table if not exists items (id integer primary key, title text, url text, likes integer, userid integer)")
	checkErr(err)
	defer statement.Close()

	statement.Exec()

	router := mux.NewRouter()
	router.Headers("Content-Type", "application/json",
		"X-Requested-With", "XMLHttpRequest")

	router.HandleFunc("/foods", getFoods).Methods("GET")
	router.HandleFunc("/food/{id}", getFood).Methods("GET")
	router.HandleFunc("/foods", addFood).Methods("POST")
	router.HandleFunc("/foods", updateFood).Methods("PUT")
	router.HandleFunc("/food/{id}", removeFood).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getFoods(w http.ResponseWriter, r *http.Request) {
	var food Food
	//assign an empty slice every time method is called
	foods = []Food{}

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

func getFood(w http.ResponseWriter, r *http.Request) {
	//var food Food

	food := Food{}
	params := mux.Vars(r)

	rows := database.QueryRow("select * from items where id=?;", params["id"])

	err := rows.Scan(&food.ID, &food.Title, &food.URL, &food.Likes, &food.UserID)
	checkErr(err)
	// handle if no entry is found
	json.NewEncoder(w).Encode(food)
}

func addFood(w http.ResponseWriter, r *http.Request) {
	//create variable type food
	var food Food

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

func updateFood(w http.ResponseWriter, r *http.Request) {
	var food Food
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

func removeFood(w http.ResponseWriter, r *http.Request) {
	//food := Food{}

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

func checkErr(err error) {
	if err != nil {
		log.Println("Yup it returned an error", err)
		//log.Fatal(err)
	}
}
