package main

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("MySecretKey")

func main() {

	// Create a JWT object with claims (id, name, and expiration time)
	jwtObject := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":   "123",
		"name": "Ali",
		"exp":  time.Now().Add(time.Hour * 24).Unix(), // token will expire after 24 hours
	})

	// Create a token string from the JWT object using the secret key
	token, err := jwtObject.SignedString(SecretKey)
	if err != nil {
		fmt.Println("Error with creating a token!")
		panic(err)
	}

	// Parse the token string and verify it using the secret key
	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		fmt.Println("Error while parsing the token!")
		panic(err)
	}

	// If the token is valid, get the claims (payload information) from it
	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if ok && parsedToken.Valid {
		fmt.Println("The user's info:")
		fmt.Println("Id:   ", claims["id"])
		fmt.Println("Name: ", claims["name"])
		fmt.Println("Exp:  ", claims["exp"])
	} else {
		fmt.Println("Error with parsing!!! ")
	}
}
