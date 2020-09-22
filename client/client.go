package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app  = kingpin.New("client", "send messages to a server")
	host = app.Flag("host", "host to connect to").Required().String()
	send = app.Flag("message", "send a request to server").Required().Strings()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))
	message := strings.Join(*send, " ")
	reader := strings.NewReader(message)
	resp, err := http.Post(*host+"/echo", "", reader)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	serverOutput := bufio.NewScanner(resp.Body)
	for serverOutput.Scan() {
		if serverOutput.Text() != message {
			fmt.Println("the server did not echo your request, oops!, got: ")
		}
		fmt.Println(serverOutput.Text())
	}
}
