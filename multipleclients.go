package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"sync"

	"gopkg.in/alecthomas/kingpin.v2"
)

//testing concurrency

func main() {
	var wg sync.WaitGroup
	kingpin.MustParse(app.Parse(os.Args[1:]))
	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go worker(&wg, *token)
	}
	wg.Wait()
}

var app = kingpin.New("clients", "connect to server")
var token = app.Flag("token", "authentication token").Required().String()

func worker(wg *sync.WaitGroup, t string) {
	defer wg.Done()
	requestMessage := "hello test"
	reader := strings.NewReader(requestMessage)

	req, err := http.NewRequest("POST", "https://localhost:8000/echo", reader)
	if err != nil {
		fmt.Println("could not resolve host: ", err)
	}
	req.Header.Add("Token", t)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("request failed: ", err)
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Not OK HTTP status:", resp.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("could not read body: ", err)
	}

	serverOutput := string(bodyBytes)
	if serverOutput != requestMessage {
		fmt.Println(" did not echo your request, oops! got:")
	}
	fmt.Println(serverOutput)
	resp.Body.Close()
}
