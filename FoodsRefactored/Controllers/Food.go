package controllers

import (
	models "FoodsRefactored/Models"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
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

func checkErr(err error) {
	if err != nil {
		log.Println("Yup it returned an error", err)
		//log.Fatal(err)
	}
}
