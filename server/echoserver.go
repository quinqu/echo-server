package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

var signingKey = []byte("secret")

func echo(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	newStr := buf.String()

	fmt.Fprintf(w, newStr)
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			token, err := jwt.Parse(r.Header["Token"][0], func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There was an error")
				}
				return signingKey, nil
			})

			if err != nil {
				fmt.Fprintf(w, err.Error())
			}

			if token.Valid {
				fmt.Println("authorized")
				endpoint(w, r)
			}
		} else {
			fmt.Fprintf(w, "Not authorized")
		}
	})

}

func main() {
	//JWT in response header
	listener, err := net.Listen("tcp", ":8000")

	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/echo", isAuthorized(echo))
	log.Println("Listening on localhost:8080")
	err = http.Serve(listener, nil)

	if err != nil {
		log.Fatal(err)

	}

}
