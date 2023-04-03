package authentication

import "time"

type RegisterData struct {
	Username  string    `json:"Username"`
	Email     string    `json:"Email"`
	Password  string    `json:"Password"`
	Gender    string    `json:"Gender"`
	Birthdate time.Time `json:"Birthdate"`
}
