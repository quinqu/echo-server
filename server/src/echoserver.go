package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

var authToken string

func echo(w http.ResponseWriter, r *http.Request) {
	var response string
	out, err := ioutil.ReadAll(r.Body)
	response = string(out)
	if err != nil {
		response = fmt.Sprintf("there was error reading request: " + err.Error())
	}
	fmt.Fprintf(w, response)

}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			if r.Header["Token"][0] == authToken {
				w.WriteHeader(http.StatusOK)
				fmt.Println("authorized")
				endpoint(w, r)
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				fmt.Fprintf(w, "not authorized")
			}

		} else {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, "token missing, unauthorized")
		}
	})

}

func generateToken(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

var (
	app  = kingpin.New("echoserver", "echo server will echo a clients request")
	port = app.Flag("port", "port to bind to").Required().String()
)

func main() {
	token, err := generateToken(20)
	if err != nil {
		log.Fatal("unable to generate token", err)
	}
	authToken = token
	fmt.Println("Authentication token: " + authToken)

	kingpin.MustParse(app.Parse(os.Args[1:]))
	serverPort := *port
	http.HandleFunc("/echo", isAuthorized(echo))
	log.Fatal(http.ListenAndServeTLS(":"+serverPort, "localhost.crt", "localhost.key", nil))
}
