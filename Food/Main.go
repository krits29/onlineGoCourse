package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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
	json.NewEncoder(w).Encode(foods)
}

func getFood(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	i, _ := strconv.Atoi(params["id"])

	for _, food := range foods {
		if food.ID == i {
			json.NewEncoder(w).Encode(food)
		}
	}

}

func addFood(w http.ResponseWriter, r *http.Request) {
	//log.Println("Add food")
	var food Food

	json.NewDecoder(r.Body).Decode(&food)

	foods = append(foods, food) //adding on

	json.NewEncoder(w).Encode(foods)
}

func updateFood(w http.ResponseWriter, r *http.Request) {
	log.Println("Update food")

	var food Food

	json.NewDecoder(r.Body).Decode(&food)

	//looking for a value in the array that matches with the sae value given

	for i, f := range foods {
		if f.ID == food.ID {
			foods[i] = food
		}
	}

	json.NewEncoder(w).Encode(foods)
}

func removeFood(w http.ResponseWriter, r *http.Request) {
	log.Println("Remove book")
}
