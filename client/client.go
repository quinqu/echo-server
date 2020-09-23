package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app   = kingpin.New("client", "send messages to a server")
	host  = app.Flag("host", "host to connect to").Required().String()
	send  = app.Flag("message", "send a request to server").Required().Strings()
	token = app.Flag("token", "authentication token").Required().String()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))
	message := strings.Join(*send, " ")
	reader := strings.NewReader(message)

	req, err := http.NewRequest("POST", *host+"/echo", reader)

	// add token header to the req
	req.Header.Add("Token", *token)

	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	serverOutput := string(bodyBytes)
	if serverOutput != message {
		fmt.Println(*host + " did not echo your request, oops! got:")
	}
	fmt.Println(serverOutput)
}
