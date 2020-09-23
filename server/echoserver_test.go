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
	handler := http.HandlerFunc(echo)
	handler.ServeHTTP(rr, req)
	bodyBytes, err := ioutil.ReadAll(rr.Body)
	response := string(bodyBytes)
	if response != message {
		t.Errorf("response err: got %v, want: %v", response, message)
	}
}