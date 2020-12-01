package mock

import "testing"

var userExistsMock func(email string) bool

type preCheckMock struct{}

func (u preCheckMock) userExists(email string) bool {
	return userExistsMock(email)
}

func TestRegisterUser(t *testing.T) {
	user := User{
		Name:     "Alexandar Kunev",
		Email:    "aalexandy@gmail.com",
		UserName: "kompir",
	}
	regPreCond := preCheckMock{}
	userExistsMock = func(email string) bool {
		return false
	}
	err := RegisterUser(user, regPreCond)
	if err != nil {
		t.Fatal(err)
	}
	userExistsMock = func(email string) bool {
		return true
	}
	err = RegisterUser(user, regPreCond)
	if err == nil {
		t.Error("Expected Register User to throw and error got nil")
	}
}
