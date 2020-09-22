package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"
)

// var (
// 	app = kingpin.New("echo", " connect to the echo server, make a user specified request, and receive the server’s response.")

// 	send    = app.Command("send", "send a request to server")
// 	message = send.Arg("message", "a (message) request to send to server").Required().Strings()
// )

func main() {

	input := bufio.NewScanner(os.Stdin)
	var buf bytes.Buffer
	
	for input.Scan() {
		buf.WriteString(input.Text())

		resp, err := http.Post("http://localhost:8000/echo", "", &buf)

		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()

		input := bufio.NewScanner(resp.Body)

		for input.Scan() {
			fmt.Println(input.Text())
		}
		if err := input.Err(); err != nil {
			log.Fatal(err)
		}
	}
}