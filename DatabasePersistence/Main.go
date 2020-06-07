package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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

func main() {
	router := mux.NewRouter()
	router.Headers("Content-Type", "application/json",
		"X-Requested-With", "XMLHttpRequest")

	foods = append(foods, Food{1, "Pasta", "https://pasta.com", 0, 0}, Food{2, "Soup", "sdfs", 0, 0}, Food{3, "Pizza", "https://pizza.com", 0, 0})

	router.HandleFunc("/foods", getFoods).Methods("GET")
	router.HandleFunc("/food/{id}", getFood).Methods("GET")
	router.HandleFunc("/foods", addFood).Methods("POST")
	router.HandleFunc("/foods", updateFood).Methods("PUT")
	router.HandleFunc("/food/{id}", removeFood).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}

func getFoods(w http.ResponseWriter, r *http.Request) {
}

func getFood(w http.ResponseWriter, r *http.Request) {
}

func addFood(w http.ResponseWriter, r *http.Request) {
}

func updateFood(w http.ResponseWriter, r *http.Request) {
}

func removeFood(w http.ResponseWriter, r *http.Request) {
}
