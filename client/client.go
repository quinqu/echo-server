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
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	serverOutput := string(bodyBytes)
	fmt.Println(serverOutput)
}
