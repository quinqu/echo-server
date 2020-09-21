package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
)


func (l *Listener) Echo(message string, reply *string) error {
	*reply = message
	return nil
}

type Listener int

func main() {
	var l = new(Listener)
	err := rpc.Register(l) // Register connection to be listened to
	if err != nil {
		log.Fatal(err)
	}

	// HandleHTTP registers an HTTP handler for RPC messages on rpcPath,
	// and a debugging handler on debugPath. It is still necessary to invoke http.Serve(),
	// typically in a go statement.
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":8000") // open the connection

	if err != nil {
		log.Fatal(err)
	}

	// Serve accepts incoming HTTP connections on the listener l,
	// creating a new service goroutine for each. The service goroutines
	// read requests and then call handler to reply to them.

	err = http.Serve(listener, nil)
	if err != nil {
		log.Fatal(err)

	}

}
