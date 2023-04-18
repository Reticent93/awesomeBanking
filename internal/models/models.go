package models

type User struct {
	ID int `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Password string `json:"password"`
}


type Account struct {
	ID int `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
	Balance float64 `json:"balance"`
	UserId int `json:"user_id"`
}

