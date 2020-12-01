package basic

import (
	"fmt"
	"go-mock-tests/db"
	"log"
)

// User struct
type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
}

// Registering uniq user
func RegisterUser(user User) error {
	// check if user is in db
	found := db.UserExists(user.Email)
	if found {
		return fmt.Errorf("email '%s' already registered", user.Email)
	}
	log.Println(user)
	return nil
}
