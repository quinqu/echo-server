package main

import (
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
)

func echo(w http.ResponseWriter, r *http.Request) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	newStr := buf.String()

	fmt.Fprintf(w, newStr)
}

type Listener int

func main() {

	listener, err := net.Listen("tcp", ":8000") // open the connection
	if err != nil {
		log.Fatal(err)
	}

	// Serve accepts incoming HTTP connections on the listener,
	// creating a new service goroutine for each. The service goroutines
	// read requests and then call handler to reply to them.

	http.HandleFunc("/echo", echo)

	log.Println("Listening on localhost:8080")

	err = http.Serve(listener, nil)

	if err != nil {
		log.Fatal(err)

	}

}
