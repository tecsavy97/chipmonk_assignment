package helpers

import (
	"chipmonk_assignment/models"
	"encoding/base64"
	"errors"
)

// GenerateToken - generate a token as per the string provided
func GenerateToken(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

// ValidateToken - validates the token sent in request header
func ValidateToken(s string) error {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return err
	}
	userMap := models.InitMap()
	user, ok := userMap[string(data)]
	if !ok && user.Username != string(data) {
		return errors.New("Please enter valid token")
	}
	return nil
}
