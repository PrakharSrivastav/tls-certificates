package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	// configure the http request
	r, err := http.NewRequest(http.MethodGet, "http://localhost:8080/server", nil)
	if err != nil {
		log.Fatalf("request failed : %v", err)
	}

	// initialize the http client
	c := http.Client{
		Timeout:   time.Second * 5,
		Transport: &http.Transport{IdleConnTimeout: 10 * time.Second},
	}

	// make the request
	data, err := callServer(c, r)
	if err != nil {
		log.Fatal(err)
	}
	print(data)
}

func callServer(c http.Client, r *http.Request) (string, error) {
	response, err := c.Do(r)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	// print the data
	return string(data), nil
}
