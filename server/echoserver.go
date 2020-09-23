package main

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)


func echo(w http.ResponseWriter, r *http.Request) {
	var response string
	out, err := ioutil.ReadAll(r.Body)
	response = string(out)
	if err != nil {
		response = fmt.Sprintf("there was error: " + err.Error())
	}
	fmt.Fprintf(w, response)
}

func isAuthorized(endpoint func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header["Token"] != nil {
			if r.Header["Token"][0] == authToken {
				fmt.Println("authorized")
				endpoint(w, r)
			}
				
		} else {
			fmt.Fprintf(w, "not authorized")
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
var authToken string

func main() {

	token, err := generateToken(20)
	if err != nil {
		log.Fatal("unable to generate token", err)
	}
	authToken = token
	fmt.Println(authToken)
	kingpin.MustParse(app.Parse(os.Args[1:]))
	serverPort := *port
	listener, err := net.Listen("tcp", ":"+serverPort)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/echo", isAuthorized(echo))
	log.Println("Listening on localhost:" + serverPort)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal(err)
	}
}
