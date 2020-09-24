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

## Quickstart: 

```
$ git clone https://github.com/quinqu/echo-server.git 
$ cd echo-server 
```

To avoid using `./binary-name` to execute the binary, copy the binary file to your `/usr/local/bin/` path.

Echo server start: 
*creates server that serves traffic over https* 
  
```
$ cd server

// create the binary
$ go build echoserver.go 

// input port to bind to
$ echoserver --port=8000

// upon initialization, server will print out an auth token
Authentication token:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx
```

Client request: 

```
$ cd client 

// create the binary
$ go build client.go 

// obtain token from server output
$ client send --host=https://localhost:8000 --message="hello world" --token=<TOKEN>
```

## Server:
- Packages:
    - https://golang.org/pkg/net/http/ 
    - https://github.com/alecthomas/kingpin
- Upon initialization of the server, a port to bind to is required 
- Encryption using TLS to secure traffic through the network 
- Testing 
    - Testing the output of the echo endpoint
        - The test port is defaulted to 8000, start server on that port (follow server quickstart instructions above)
        - Run:  `go test echoserver_test.go`
### Authentication: 
- Generate authentication token on server startup using the [cryto/rand](https://golang.org/pkg/crypto/rand/) package 

### Concurrency: 
- Using `http.ListenAndServeTLS` function from the [net/http](https://golang.org/pkg/net/http/) which returns `http.Serve` and this method creates a new service goroutine for each incoming connection, [source code](https://golang.org/src/net/http/server.go)
- Verify concurrency with `multipleclients.go` in the project root 




## Client: 
- Packages: 
    - https://github.com/alecthomas/kingpin     
    - https://golang.org/pkg/net/http/ 
- Response is verified by simply  comparing the request and response 



