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

#### Notes: 
I have no plans on implementing authentication or authorization yet (depends on how much time I have) 

### Implementation: 

#### Concurrency: 
I will use http.Serve (in the net/http library) which creates a new service goroutine for each incoming HTTP connection https://golang.org/src/net/http/server.go


#### Server:
- Packages:
    - https://golang.org/pkg/net/http/
    - https://golang.org/pkg/net/rpc/ 
    - API mechanism: HTTP
    - Server listening on port 8000 for incoming client requests


#### Client: 

- Packages: 
    - https://github.com/alecthomas/kingpin for CLI 
    - https://golang.org/pkg/net/rpc/ to access echo endpoint 
    - I will verify the response is valid simply by comparing the request and response 




