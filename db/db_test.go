package db

import "testing"

func TestUserExist(t *testing.T) {
	found := UserExists("aalexandyr@gmail.com")
	if found != true {
		t.Errorf("expected true")
	}
}
