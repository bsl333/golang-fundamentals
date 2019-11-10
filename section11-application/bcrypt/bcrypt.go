package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	password := "password"

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bytes))

	// mock login
	// normally bytes would be retrieved from DB
	validPass := password
	if err := bcrypt.CompareHashAndPassword(bytes, []byte(validPass)); err != nil {
		fmt.Println(err)
		fmt.Println("invalid login info")
		return
	}

	fmt.Println("logged in - redirect")

}
