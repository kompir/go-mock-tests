package db

// simple db
var db map[string]bool

// init initialize db
func init() {
	db = make(map[string]bool)
	db["aalexandy@gmail.com"] = true
	db["kompir@example.com"] = true
}

// UserExists check if the User is registered with the provided email.
func UserExists(email string) bool {
	if _, ok := db[email]; !ok {
		return false
	}
	return true
}
