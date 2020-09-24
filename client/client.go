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
	app = kingpin.New("client", "connect to server")

	send    = app.Command("send", "send message to server")
	host    = send.Flag("host", "host to connect to").Required().String()
	message = send.Flag("message", "send a request to server").Required().Strings()
	token   = app.Flag("token", "authentication token").Required().String()
)

func main() {
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	case send.FullCommand():
		requestMessage := strings.Join(*message, " ")
		reader := strings.NewReader(requestMessage)
		endpoint := *host + "/echo"
		req, err := http.NewRequest("POST", endpoint, reader)
		if err != nil {
			fmt.Println("request could not be initiated: ", err)
			return 
		}
		req.Header.Add("Token", *token)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("response error: ", err)
			return 
		}

		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Println("Not OK HTTP status:", resp.StatusCode)
			return
		}

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("could not read body: %v", err)
		}

		serverOutput := string(bodyBytes)
		if serverOutput != requestMessage {
			fmt.Println(*host + " did not echo your request, oops! got:")
		}
		fmt.Println(serverOutput)
	}
}
