package main

import (
	"fmt"
	"io/ioutil"
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
			fmt.Println("could not resolve host: ", err)
			os.Exit(1)
			return
		}
		req.Header.Add("Token", *token)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("request failed: ", err)
			os.Exit(1)
			return
		}

		defer resp.Body.Close()
		if resp.StatusCode != http.StatusOK {
			fmt.Println("Not OK HTTP status:", resp.StatusCode)
			os.Exit(1)
			return
		}

		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("could not read body: ", err)
			os.Exit(1)
			return
		}

		serverOutput := string(bodyBytes)
		if serverOutput != requestMessage {
			fmt.Println(*host + " did not echo your request, oops! got:")
		}
		fmt.Println(serverOutput)
	default:

		fmt.Println("command not given")
	}
}
