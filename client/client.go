package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
	"time"

	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	app = kingpin.New("client", "connect to a server")

	bench    = app.Command("bench", "get latency profile")
	bHost    = bench.Flag("host", "bench host").Required().String()
	bMessage = bench.Flag("message", "message to send").Default("test test").Strings()
	bWorkers = bench.Flag("workers", "quantity of goroutines").Default("20").Int()

	send = app.Command("send", "send message to server")
	host = send.Flag("host", "host to connect to").Required().String()
	m    = send.Flag("message", "send a request to server").Required().Strings()
)

func main() {

	switch kingpin.MustParse(app.Parse(os.Args[1:])) {
	// Register user
	case bench.FullCommand():

		// start := time.Now()
		// ... operation that takes 20 milliseconds ...

		//
		var wg sync.WaitGroup

		testMessage := strings.Join(*bMessage, " ")
		for i := 0; i < *bWorkers; i++ {
			wg.Add(1)
			go worker(*bHost, testMessage, &wg)
		}
		wg.Wait()

	case send.FullCommand():
		message := strings.Join(*m, " ")
		reader := strings.NewReader(message)
		resp, err := http.Post(*host+"/echo", "", reader)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		serverOutput := string(bodyBytes)
		if serverOutput != message {
			fmt.Println(*host + " did not echo your request, oops! got:")
		}
		fmt.Println(serverOutput)

	}

}

// RequestsOriginated is amount of reuqests originated
var RequestsOriginated int
var RequestsFailed int

func worker(host, message string, wg *sync.WaitGroup) {
	defer wg.Done()


	for i := 0; i < 20; i++ {
		start := time.Now()
		reader := strings.NewReader(message)
		res, err := http.Post(host + "/echo", "", reader)
		if err != nil {
			log.Fatal(err)
			RequestsFailed += 1
		}
		RequestsOriginated += 1
		defer res.Body.Close()
		bodyBytes, err := ioutil.ReadAll(res.Body)
		serverOutput := string(bodyBytes)
		log.Print(serverOutput, ": ")
		t := time.Now()
		elapsed := t.Sub(start)
		log.Println(elapsed)
	}

}

// for _, quantile := range []float64{25, 50, 75, 90, 95, 99, 100} {
// 	fmt.Fprintf(t, "%v\t%v ms\n", quantile, result.Histogram.ValueAtQuantile(quantile))
// }
