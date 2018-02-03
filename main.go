package main

//entry point for money-mon backend service

import (
	"net/http"

	"github.com/IssueSquare/money-mom/cmd/user"
)

func main() {
	http.HandleFunc("/userRegist", user.Register)
	http.HandleFunc("/userLogin", user.Login)
	http.ListenAndServe(":8080", nil)
}
