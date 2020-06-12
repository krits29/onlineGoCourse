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

	// router.HandleFunc("/foods", controller.GetFoods(database)).Methods("GET")
	// router.HandleFunc("/food/{id}", controller.GetFood(database)).Methods("GET")
	// router.HandleFunc("/foods", controller.AddFood(database)).Methods("POST")
	// router.HandleFunc("/foods", controller.UpdateFood(database)).Methods("PUT")
	// router.HandleFunc("/food/{id}", controller.RemoveFood(database)).Methods("DELETE")

	router.HandleFunc("/user/{id}/foods", controller.GetFoods(database)).Methods("GET")
	router.HandleFunc("/user/{id}/food/{id}", controller.GetFood(database)).Methods("GET")
	router.HandleFunc("/user/{id}/foods", controller.AddFood(database)).Methods("POST")
	router.HandleFunc("/user/{id}/foods", controller.UpdateFood(database)).Methods("PUT")
	router.HandleFunc("/user/{id}/food/{id}", controller.RemoveFood(database)).Methods("DELETE")

	router.HandleFunc("/users", controller.GetUsers(database)).Methods("GET")
	router.HandleFunc("/register", controller.Register(database)).Methods("POST")
	router.HandleFunc("/login", controller.Login(database)).Methods("POST")

	router.HandleFunc("/user/{id}", controller.GetUser(database)).Methods("GET")

	//router.HandleFUnc("/logout", ????).Methods("POST")

	fmt.Println("Server is running")

	log.Fatal(http.ListenAndServe(":8000", router))

	//todo
	//no duplicate
	//password should be encrypted
	//getUser should not contain password

}
