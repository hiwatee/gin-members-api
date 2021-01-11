package service

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

// HashAndSalt ...
func HashAndSalt(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}
