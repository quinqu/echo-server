package main

import (
	"fmt"
	"log"
	"net/rpc"
	"os"
	"strings"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("echo", " connect to the echo server, make a user specified request, and receive the server’s response.")

	send    = app.Command("send", "send a request to server")
	message = send.Arg("message", "a (message) request to send to server").Required().Strings()
)

func main() {
	var reply string
	var client *rpc.Client
	var err error
	switch kingpin.MustParse(app.Parse(os.Args[1:])) {

	case send.FullCommand():
		client, err = rpc.DialHTTP("tcp", "localhost:8000")

		if err != nil {
			log.Fatal("dialing: ", err)
		}
		echo := strings.Join(*message, " ")
		client.Call("Listener.Echo", echo, &reply)
		fmt.Println(reply)
	}

	//DialHTTP connects to an HTTP RPC server at the specified network
	//address listening on the default HTTP RPC path.

}
