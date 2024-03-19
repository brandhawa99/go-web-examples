package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func main() {
	password := "secret"
	hash, _ := HashPassword(password)
	match := CheckPasswordHash(password, hash)

	fmt.Println("password: ", password)
	fmt.Println("Hash: ", hash)
	fmt.Println()
	fmt.Println("Match:		", match, "")

	password2 := hash
	hash2, _ := HashPassword(password2)
	match2 := CheckPasswordHash(password2, hash2)

	fmt.Println("password", hash)
	fmt.Println("hashed hash: ", hash2)
	fmt.Println("Match:   ", match2)

}
