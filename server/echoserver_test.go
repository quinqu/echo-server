package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestEcho(t *testing.T) {
	message := "hello test"
	reader := strings.NewReader(message)
	req, err := http.NewRequest("POST", "/echo", reader)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	
	echo(rr, req)
	bodyBytes, err := ioutil.ReadAll(rr.Body)
	if err != nil {
		t.Errorf("error reading data: %v", err)
	}
	response := string(bodyBytes)
	if response != message {
		t.Errorf("response does not match request: got %v, want: %v", response, message)
	}
}
