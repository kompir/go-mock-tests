package mock

import (
	"fmt"
	"go-mock-tests/db"
	"log"
)

type User struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	UserName string `json:"user_name"`
}

type RegistrationPreChecker interface {
	userExists(string) bool
}

type regPreCheck struct{}

func (r regPreCheck) userExists(email string) bool {
	return db.UserExists(email)
}

func NewRegistrationPreChecker() RegistrationPreChecker {
	return regPreCheck{}
}

func RegisterUser(user User, regPreCond RegistrationPreChecker) error {
	found := regPreCond.userExists(user.Email)
	if found {
		return fmt.Errorf("email '%s' already registered", user.Email)
	}
	log.Println(user)
	return nil
}
