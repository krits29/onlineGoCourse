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
	json.NewEncoder(w).Encode(foods)
}

func getFood(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	//need to convert because the number is interpreted as a string
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

	//final operation needed when making changes (kind of like a return)
	json.NewEncoder(w).Encode(foods)
}

func updateFood(w http.ResponseWriter, r *http.Request) {
	log.Println("Update food") //call debugging purpose

	var food Food

	json.NewDecoder(r.Body).Decode(&food)

	//looking for a value in the array that matches with the same value given

	for i, foood := range foods {
		if foood.ID == food.ID {
			foods[i] = food
		}
	}

	//final operation needed when making changes (kind of like a return)
	json.NewEncoder(w).Encode(foods)
}

func removeFood(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) //returns all the parameters in the request

	//need to convert because the number is interpreted as a string
	id, _ := strconv.Atoi(params["id"])

	// var newFood []Food

	// for i, food := range foods {
	// 	if food.ID != id {
	// 		newFood = append(newFood, foods[i])
	// 	}
	// }

	// foods = newFood

	for i, food := range foods {
		if food.ID == id {
			foods = append(foods[:i], foods[i+1:]...)
		}
	}

	json.NewEncoder(w).Encode(foods)
}
