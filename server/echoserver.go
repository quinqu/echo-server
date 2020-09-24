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
	sem <- struct{}{}
	defer func() { <-sem }()
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
	sem = make(chan struct{}, 20)
)


func main() {
	
	token, err := generateToken(20)
	if err != nil {
		log.Fatal("unable to generate token", err)
	}
	authToken = token
	fmt.Println(authToken)
	kingpin.MustParse(app.Parse(os.Args[1:]))
	serverPort := *port

	http.HandleFunc("/echo", echo)//isAuthorized(echo))

	// ListenAndServe invokes Serve, which calls Accept in a loop. 
	// Accept is similar to the socket accept syscall, creating a new 
	// connection whenever a new client is accepted; the connection is then 
	// used for interacting with the client. This connection is type conn in the 
	// server.go file, a private type with server state for each client connection. 
	// Its serve method actually serves a connection, pretty much as you would expect; 
	// it reads some data from the client and invokes the user-suppied handler depending on the path.

	log.Fatal(http.ListenAndServeTLS(":"+serverPort, "localhost.crt", "localhost.key", nil))

}
