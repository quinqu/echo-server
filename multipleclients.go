package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"sync"
)

func main() {
	fmt.Println("go func")
	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
}


func worker(wg *sync.WaitGroup) {
	defer wg.Done()

	reader := strings.NewReader("hello world")
	res, err := http.Post("https://localhost:8000/echo", "", reader)
	if err != nil {
		log.Fatal(err)
	}

	 
	bodyBytes, err := ioutil.ReadAll(res.Body)
	serverOutput := string(bodyBytes)
	log.Println(serverOutput)
	res.Body.Close()
}