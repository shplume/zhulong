package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// Generate a salt hash.
func GenerateSalt(password *string) (string, error) {
	hex := []byte(*password)
	hashedPassword, err := bcrypt.GenerateFromPassword(hex, 10)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// Compare a salt hash.
func CompareSalt(digest []byte, password *string) bool {
	hex := []byte(*password)
	if err := bcrypt.CompareHashAndPassword(digest, hex); err == nil {
		return true
	}
	return false
}
