package models

//User is
type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}