package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"gopkg.in/alecthomas/kingpin.v2"
)

var signingKey = []byte("secret") //from env variables os.Get("MY_JWT_TOKEN")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	tokenString, err := token.SignedString(signingKey)

	if err != nil {
		fmt.Errorf("Something went wrong: %s", err)
		return "", err
	}
	return tokenString, nil
}

var (
	app  = kingpin.New("echo", " connect to the echo server, make a user specified request")
	port = app.Flag("port", "server to connect to ex: '--port=8000'").Required().String()
	send = app.Flag("message", "send a request to server").Required().Strings()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))
	tokenString, err := GenerateJWT()

	var buf bytes.Buffer
	message := strings.Join(*send, " ")
	url := "http://localhost:" + *port + "/echo"

	buf.WriteString(message)

	// Create a new request using http
	req, err := http.NewRequest("POST", url, &buf)

	// add authorization header to the req
	req.Header.Add("Token", tokenString)

	// Send req using http Client
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	serverOutput := bufio.NewScanner(resp.Body)

	for serverOutput.Scan() {
		if serverOutput.Text() != message {
			fmt.Println("The server did not echo your request, oops!, Got: ")

		}
		fmt.Println(serverOutput.Text())
	}

}
