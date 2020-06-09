package utils

import (
	"FoodsRefactored/models"
	"encoding/json"
	"log"
	"net/http"
)

func sendError(w http.ResponseWriter, status int, err models.Error) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(err)
}

func sendSuccess(w http.ResponseWriter, data interface{}) {
	json.NewEncoder(w).Encode(data)
}

func CheckErr(err error) {
	if err != nil {
		log.Println("Yup it returned an error", err)
		//log.Fatal(err)
	}
}
