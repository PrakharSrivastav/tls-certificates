package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	r, err := http.NewRequest(http.MethodGet, "http://localhost:8080/server", nil)
	if err != nil {
		log.Fatalf("request failed : %v", err)
	}

	c := http.Client{
		Timeout:   time.Second * 5,
		Transport: &http.Transport{IdleConnTimeout: 10 * time.Second},
	}

	response, err := c.Do(r)
	if err != nil {
		log.Fatalf("response failed : %v", err)
	}
	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("body read failed :  %v", err)
	}

	log.Println(string(data))
}
