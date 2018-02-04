package main

//entry point for money-mon backend service

import (
	"net/http"

	"github.com/IssueSquare/money-mom/cmd/user"
	"github.com/IssueSquare/money-mom/pkg/auth"
)

type basicCredential struct {
	clientID string
	password string
}

func main() {

	//TODO load Basic account from DB
	bc := &basicCredential{
		clientID: "ios-x02",
		password: "testpass",
	}

	//User api group should secure with Basic token.
	http.HandleFunc("/userRegister", auth.BasicAuth(user.Register, []byte(bc.clientID), []byte(bc.password)))
	http.HandleFunc("/userLogin", auth.BasicAuth(user.Login, []byte(bc.clientID), []byte(bc.password)))

	http.ListenAndServe(":8080", nil)
}
