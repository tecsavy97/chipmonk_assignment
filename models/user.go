package models

import "sync"

var (
	// userMap -  to store the user data
	userMap map[string]User
	// so - used as single initialisation factor needed for the userMap
	so sync.Once
)

// Users - Data structure to store user data
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IsActive bool   `json:"isActive,omitempty"`
}

// InitMap -  to intialise the userMap once
func InitMap() map[string]User {
	so.Do(func() {
		userMap = make(map[string]User)
	})
	return userMap
}

// NewUser - returns a new Users empty str
func NewUser() User {
	return User{}
}
