## Overview: 

Implement a prototype concurrent “echo server” and client. The server should return the client’s request as the server response.

## Requirements

### Server
The echo server should provide a single endpoint that returns the verbatim request issued to it. Any API mechanism that works for the task and is familiar to you is OK: Raw Sockets, HTTP, HTTPS/JSON, GRPC, or anything else that can guarantee reliable client-server communication.
The echo server must support concurrent requests from different clients.

### Client
The client command should be able to connect to the echo server, make a user specified request, and receive the server’s response.
The client should verify the response is equivalent to the request, and print the response.
The client should provide cli flags or arguments to control its behavior.


## Server:
- Packages:
    - https://golang.org/pkg/net/http/ 
    - https://github.com/alecthomas/kingpin
- Upon initialization of the server, a port to bind to is required 
- Encryption using TLS to secure traffic through the network 

### Authentication: 
- Generate authentication token on server startup using the [cryto/rand](https://golang.org/pkg/crypto/rand/) package 

### Concurrency: 
- Using `http.ListenAndServeTLS` function from the [net/http](https://golang.org/pkg/net/http/) which is built on top of [http.Serve](https://golang.org/src/net/http/server.go) and this method creates a new service goroutine for each incoming connection 

Start an echo server example: 

 `$ echoserver --port=8000`



## Client: 
- Packages: 
    - https://github.com/alecthomas/kingpin     
    - https://golang.org/pkg/net/http/ 
- Response is verified by simply  comparing the request and response 

Connect and request example: 

  `$ client send --host=https:localhost:8000 --message="echo me" --token=<TOKEN>`



