package auth

import (
	"encoding/base64"
	"fmt"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

//BasicAuth a way of verify client-side
func BasicAuth(auth []string) error {

	if len(auth) != 2 || auth[0] != "Basic" {
		return fmt.Errorf("error")
		//check the token
	}
	token, _ := base64.StdEncoding.DecodeString(auth[1])
	pair := strings.SplitN(string(token), ":", 2)

	if len(pair) != 2 || !ValidateBasicToken(pair[0], pair[1]) {
		return fmt.Errorf("error")
	}
	return nil
}

//ValidateJWT a way of verify JWT
func ValidateJWT(auth []string) error {
	if len(auth) != 2 || auth[0] != "Bearer" {
		return fmt.Errorf("error")
		//check the token
	}
	_, err := jwt.ParseWithClaims(auth[1], nil, func(token *jwt.Token) (interface{}, error) {
		return []byte("sharekey"), nil
	})

	if err != nil {
		panic(err)
	}
	return nil
}

//ValidateBasicToken used to verify
func ValidateBasicToken(username, password string) bool {
	//todo should put in db
	if username == "money-ios-client" && password == "xdfg212" {
		return true
	}
	return false
}
