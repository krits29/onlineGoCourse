package main

import (
	"database/sql"
	"fmt"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/mattn/go-sqlite3"

	"FoodsRefactored/controllers"
	"FoodsRefactored/driver"
	"FoodsRefactored/models"
)

var foods []models.Food

var database *sql.DB

func main() {

	database = driver.ConnectDB()
	defer database.Close()
	controller := controllers.Controller{}

	router := mux.NewRouter()
	router.Headers("Content-Type", "application/json",
		"X-Requested-With", "XMLHttpRequest")

	router.HandleFunc("/foods", controller.GetFoods(database)).Methods("GET")
	router.HandleFunc("/food/{id}", controller.GetFood(database)).Methods("GET")
	router.HandleFunc("/foods", controller.AddFood(database)).Methods("POST")
	router.HandleFunc("/foods", controller.UpdateFood(database)).Methods("PUT")
	router.HandleFunc("/food/{id}", controller.RemoveFood(database)).Methods("DELETE")
	//router.HandleFUnc("/register", ????).Methods("POST")
	//router.HandleFUnc("/login", ????).Methods("POST")
	//router.HandleFUnc("/logout", ????).Methods("POST")

	fmt.Println("Server is running")

	log.Fatal(http.ListenAndServe(":8000", router))
}
