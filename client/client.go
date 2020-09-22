package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("client", "send a request to a server!")

	port = app.Flag("port", "port to connect to ex: --port=8000").Required().String()
	send = app.Flag("message", "send a message to the server").Required().Strings()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))

	var buf bytes.Buffer

	message := strings.Join(*send, " ")
	buf.WriteString(message)
	host := "http://localhost:" + *port + "/echo"
	resp, err := http.Post(host, "", &buf)

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	serverOutput := bufio.NewScanner(resp.Body)

	for serverOutput.Scan() {
		// compare response with request
		if serverOutput.Text() != message {
			fmt.Println("response is not the same as your message, oops!, got: ")
		}
		fmt.Println(serverOutput.Text())
	}

}
