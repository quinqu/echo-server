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

func main() {

	listener, err := net.Listen("tcp", ":8000") 
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/echo", echo)
	log.Println("Listening on localhost:8080")

	err = http.Serve(listener, nil)

	if err != nil {
		log.Fatal(err)

	}

}
