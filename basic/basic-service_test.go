package basic

import "testing"

func TestCheckUserExist(t *testing.T) {
	user := User{
		Name:     "Alexandar Kunev",
		Email:    "aalexandy@gmail.com",
		UserName: "kompir",
	}
	err := RegisterUser(user)
	if err == nil {
		t.Error("Expected Register User to throw and error got nil")
	}
}
