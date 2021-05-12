package helpers

import "golang.org/x/crypto/bcrypt"

// HashPassword -  hash password to store in db
func HashPassword(p string) string {
	b, _ := bcrypt.GenerateFromPassword([]byte(p), bcrypt.DefaultCost)
	return string(b)
}

// ValidatePassword - validates the hashed password
func ValidatePassword(h, p string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(h), []byte(p))
	return err == nil
}
