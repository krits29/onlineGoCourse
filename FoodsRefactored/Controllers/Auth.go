package controllers

import (
	"FoodsRefactored/models"
	"database/sql"
	"encoding/json"
	"strings"

	// 	"log"
	"io/ioutil"
	"net/http"
	// 	"github.com/gorilla/mux"
)

func (c Controller) Validate(action string) error {
	var user models.User
	var errors models.Error
	switch strings.ToLower(action) {
	case "login":
		if user.Username == "" {
			errors.Message = ("Email is required")
			//return
		}
		if user.Password == "" {
			errors.Message = ("Password is required")
			//return
		}
		return nil
	default: // this is for creating a user, where all fields are required
		if user.Username == "" {
			errors.Message = ("Email is required")
			//return
		}
		if user.Password == "" {
			errors.Message = ("Password is required")
			//return
		}

		//checking valid email format
		//not necessary for now
		// if err := checkmail.ValidateFormat(user.Username); err != nil {
		// 	errors.Message = ("Invalid Email")
		// 	//return
		// }
		return nil
	}
}

func (c Controller) Login(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var resp = map[string]interface{}{"status": "success", "message": "logged in"}

		user := &models.User{}
		body, err := ioutil.ReadAll(r.Body) // read user input from request
		checkErr(err)

		err = json.Unmarshal(body, &user)
		checkErr(err)

		//user.Prepare() // here strip the text of white spaces

		// err = user.Validate("login") // fields(email, password) are validated
		// checkErr(err)

		usr, err := GetUser(database) //mismatch return types and number???
		checkErr(err)

		if usr == nil { // user is not registered
			resp["status"] = "failed"
			resp["message"] = "Login failed, please signup"
			//responses.JSON(w, http.StatusBadRequest, resp)
			return
		}

		err = models.CheckPasswordHash(user.Password, usr.Password)
		if err != nil {
			resp["status"] = "failed"
			resp["message"] = "Login failed, please try again"
			responses.JSON(w, http.StatusForbidden, resp)
			return
		}
		token, err := utils.EncodeAuthToken(usr.ID)
		checkErr(err)

		resp["token"] = token
		responses.JSON(w, http.StatusOK, resp)
		return
	}
}

// //Controller is
// type Controller struct{}

// var foods []models.Food

// func (c Controller) Register(database *sql.DB) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		var food models.Food
// 		//assign an empty slice every time method is called
// 		foods = []models.Food{}

// 		rows, err := database.Query("select * from items")
// 		checkErr(err)
// 		defer rows.Close()

// 		for rows.Next() {
// 			rows.Scan(&food.ID, &food.Title, &food.URL, &food.Likes, &food.UserID)
// 			checkErr(err)

// 			foods = append(foods, food)
// 		}

// 		json.NewEncoder(w).Encode(foods)
// 	}
// }

// func checkErr(err error) {
// 	if err != nil {
// 		log.Println("Yup it returned an error", err)
// 		//log.Fatal(err)
// 	}
// }
