package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"

	"gopkg.in/alecthomas/kingpin.v2"
)

func echo(w http.ResponseWriter, r *http.Request) {
	var response string
	out, err := ioutil.ReadAll(r.Body)
	response = string(out)
	if err != nil {
		response = fmt.Sprintf("there was error: " + err.Error())
	}
	fmt.Fprintf(w, response)
}

var (
	app  = kingpin.New("echoserver", "echo server will echo a clients request")
	port = app.Flag("port", "server to connect to ex: --port=8000").Required().String()
)

func main() {
	kingpin.MustParse(app.Parse(os.Args[1:]))
	serverPort := *port
	listener, err := net.Listen("tcp", ":"+serverPort)
	if err != nil {
		log.Fatal(err)
	}
	http.HandleFunc("/echo", echo)
	log.Println("Listening on localhost:" + serverPort)
	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal(err)
	}
}
