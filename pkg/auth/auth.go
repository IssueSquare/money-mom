package auth

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

//CallBackFunc used to fit http package ineterface
type CallBackFunc func(http.ResponseWriter, *http.Request)

//BasicAuth func used to verify User relative serviecs
func BasicAuth(f CallBackFunc, user, passwd []byte) CallBackFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//fetch request header
		auth := r.Header.Get("Authorization")
		if strings.HasPrefix(auth, "Basic ") {
			//decode basic token
			payload, err := base64.StdEncoding.DecodeString(
				auth[len("Basic "):],
			)
			if err == nil {
				pair := bytes.SplitN(payload, []byte(":"), 2)
				//TODO last pair will get newline char.
				pair[1] = bytes.TrimSuffix(pair[1], []byte{'\n'})

				if len(pair) == 2 && bytes.Equal(pair[0], user) &&
					bytes.Equal(pair[1], passwd) {
					//if pass invoke callback func
					f(w, r)
					return
				}
			}
		}

		//verify faild.
		w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
		w.WriteHeader(http.StatusUnauthorized)
	}
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
