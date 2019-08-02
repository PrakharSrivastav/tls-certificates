package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

const RootCertificatePath string = "../00-certificates/minica.pem"

func main() {
	var err error
	var data string
	var r *http.Request

	// create a Certificate pool to hold one or more CA certificates
	// read minica certificate and add to the Certificate Pool
	rootCAPool := x509.NewCertPool()
	rootCA, err := ioutil.ReadFile(RootCertificatePath)
	if err != nil {
		log.Fatalf("reading cert failed : %v", err)
	}
	rootCAPool.AppendCertsFromPEM(rootCA)
	log.Println("RootCA loaded")

	c := http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout: 10 * time.Second,
			TLSClientConfig: &tls.Config{RootCAs: rootCAPool,},
		},
	}

	if r, err = http.NewRequest(http.MethodGet, "https://server-cert:8080/server", nil); err != nil {
		log.Fatalf("request failed : %v", err)
	}

	// make the request
	if data, err = callServer(c, r); err != nil {
		log.Fatal(err)
	}
	log.Println(data)
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
