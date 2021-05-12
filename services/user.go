package services

import (
	"chipmonk_assignment/helpers"
	"chipmonk_assignment/models"
	"errors"
)

// RegisterUser - Saves the data in a map and check whether the user data is previously present in map or not
func RegisterUser(u models.User) error {
	userMap := models.InitMap()
	_, ok := userMap[u.Username]
	if ok {
		return errors.New("Please Login!!")
	}
	hPass := helpers.HashPassword(u.Password)
	u.Password = hPass
	u.IsActive = false
	userMap[u.Username] = u
	return nil
}

// LoginUser - validates the user and creates a session token for that user
func LoginUser(u models.User) (string, error) {
	userMap := models.InitMap()
	data, ok := userMap[u.Username]
	if !ok {
		return "", errors.New("User not found. Please register!!")
	}
	val := helpers.ValidatePassword(data.Password, u.Password)
	if !val {
		return "", errors.New("Username/Password is incorrect")
	}
	data.IsActive = true
	userMap[data.Username] = data
	token := helpers.GenerateToken(u.Username)
	return token, nil
}

// FetchActiveUsers - get the list of all the active data in the map
func FetchActiveUsers(token string) (interface{}, error) {
	var activeUser []string
	valTokenErr := helpers.ValidateToken(token)
	if valTokenErr != nil {
		return nil, valTokenErr
	}
	userMap := models.InitMap()
	for _, v := range userMap {
		if v.IsActive {
			activeUser = append(activeUser, v.Username)
		}
	}
	return activeUser, nil
}
