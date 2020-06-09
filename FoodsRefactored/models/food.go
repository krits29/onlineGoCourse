package models

//Food is
type Food struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	URL    string `json:"url"`
	Likes  int    `json:"likes"`
	UserID int    `json:"userid"`
}