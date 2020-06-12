package controllers

import (
	"FoodsRefactored/models"
	"FoodsRefactored/utils"
	"database/sql"
	"encoding/json"
	"strings"

	// 	"log"
	"io/ioutil"
	"net/http"
	// 	"github.com/gorilla/mux"
)

//Login is
func (c Controller) Login(database *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//var resp = map[string]interface{}{"status": "success", "message": "logged in"}

		var error models.Error

		user := &models.User{}
		body, err := ioutil.ReadAll(r.Body) // read user input from request
		checkErr(err)

		err = json.Unmarshal(body, &user)
		checkErr(err)

		//user.Prepare() // here strip the text of white spaces

		err = validate("login") // fields(email, password) are validated
		checkErr(err)

		// Get this user from the database
		fetchedUser, err := GetUserFromDB(database, user)
		checkErr(err)

		//is this correct

		if fetchedUser.ID == 0 { // user is not registered
			//resp["status"] = "failed"
			//resp["message"] = "Login failed, please signup"
			//responses.JSON(w, http.StatusBadRequest, resp)

			error.Message = "Login failed, please signup"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}

		err = utils.CheckPasswordHash(user.Password, fetchedUser.Password)
		if err != nil {
			//resp["status"] = "failed"
			//resp["message"] = "Login failed, please try again"
			//responses.JSON(w, http.StatusForbidden, resp)

			error.Message = "Login failed, please try again"
			utils.SendError(w, http.StatusBadRequest, error)
			return
		}
		// user is authenticated

		token, err := utils.EncodeAuthToken(fetchedUser.ID)
		checkErr(err)

		//resp["token"] = token

		//explain??
		error.Message = token
		utils.SendError(w, http.StatusBadRequest, error)
		return
	}
}

func validate(action string) error {
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
