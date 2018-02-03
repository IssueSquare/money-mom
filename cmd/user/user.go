package user

//User model implement

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//User data struct
type User struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type Payload struct {
	Email string
	Name  string
	jwt.StandardClaims
}

//Register a user func.
func Register(rw http.ResponseWriter, req *http.Request) {
	var u User
	decoder := json.NewDecoder(req.Body)
	err := decoder.Decode(&u)

	if err != nil {
		panic(err)
	}

	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	//query persist storage system to check whether user info exist
	c := session.DB("moneymom").C("user")
	n, err := c.Find(bson.M{"email": u.Email}).Count()
	if err != nil {
		panic(err)
	}

	if n == 0 {
		err := c.Insert(&u)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(rw, "yeah welcome %s to money-mon", u.Name)
	} else {
		fmt.Fprintf(rw, "woo %s have registerd already", u.Name)
	}
}

type LoginInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//Login func.
func Login(rw http.ResponseWriter, req *http.Request) {

	decoder := json.NewDecoder(req.Body)
	var i LoginInfo
	err := decoder.Decode(&i)

	if err != nil {
		panic(err)
	}

	session, err := mgo.Dial("mongodb://localhost:27017")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	// Optional. Switch the session to a monotonic behavior.
	session.SetMode(mgo.Monotonic, true)

	//query persist storage system to check whether user info exist
	c := session.DB("moneymom").C("user")
	var u User
	err = c.Find(bson.M{"email": i.Email, "password": i.Password}).One(&u)
	if err != nil {
		fmt.Fprintf(rw, "sorry wrong email or pass")
	} else {

		// Embed User information to `token`
		token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), &Payload{Email: u.Email, Name: u.Name})
		// token -> string. Only server knows this secret (foobar).
		tokenstring, err := token.SignedString([]byte("sharekey"))
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Fprintf(rw, tokenstring)
	}
}
