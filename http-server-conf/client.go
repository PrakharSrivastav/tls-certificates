package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	c := http.Client{
		Timeout:   5 * time.Second,
		Transport: &http.Transport{IdleConnTimeout: 10 * time.Second,},
	}

	// req, err := http.NewRequest(http.MethodGet, "https://localhost:8080/server", nil)
	req, err := http.NewRequest(http.MethodGet, "https://server-cert:8080/server", nil)
	if err != nil {
		log.Fatalf("request failed : %v", err)
	}

	response, err := c.Do(req)
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

/*
req, err := http.NewRequest(http.MethodGet, "http://localhost:8080/server", nil)
Server
2019/07/25 12:49:01 http: TLS handshake error from 127.0.0.1:37304: tls: first record does not look like a TLS handshake

Client
2019/07/25 12:49:01 response failed : Get http://localhost:8080/server: net/http: HTTP/1.x transport connection broken: malformed HTTP response "\x15\x03\x01\x00\x02\x02"
exit status 1



**********************
req, err := http.NewRequest(http.MethodGet, "https://localhost:8080/server", nil)
Server
2019/07/25 12:49:35 http: TLS handshake error from 127.0.0.1:37306: remote error: tls: bad certificate

Client
2019/07/25 12:49:35 response failed : Get https://localhost:8080/server: x509: certificate is valid for server-cert, not localhost
exit status 1

*********************
req, err := http.NewRequest(http.MethodGet, "https://server-cert:8080/server", nil)
Server
2019/07/25 12:50:15 http: TLS handshake error from 127.0.0.1:37308: remote error: tls: bad certificate

Client
2019/07/25 12:50:15 response failed : Get https://server-cert:8080/server: x509: certificate signed by unknown authority
exit status 1

 */