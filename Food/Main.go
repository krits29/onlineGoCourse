package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

type Food struct {
	ID int 'json:id'
	Title string 'json:title' 
	URL string 'json:url' 
	Likes int 'json:likes'
	UserID int 'json:userid' 
}

var foods []Food

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/foods", getFoods).Methods("GET")
	router.HandleFunc("/foods/{id}", getFood).Methods("GET")
	router.HandleFunc("/foods", addFood).Methods("POST")
	router.HandleFunc("/foods", updateFood).Methods("PUT")
	router.HandleFunc("/foods/{id}", removeFood).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}