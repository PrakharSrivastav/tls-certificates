package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	RootCertificatePath string = "../certificates/minica.pem"
)

func main() {

	rootCA, err := ioutil.ReadFile(RootCertificatePath)
	if err != nil {
		log.Fatalf("reading cert failed : %v", err)
	}
	rootCAPool := x509.NewCertPool()
	rootCAPool.AppendCertsFromPEM(rootCA)
	log.Println("RootCA loaded")

	c := http.Client{
		Timeout: 5 * time.Second,
		Transport: &http.Transport{
			IdleConnTimeout: 10 * time.Second,
			TLSClientConfig: &tls.Config{
				RootCAs: rootCAPool,
			},
		},
	}

	u := url.URL{
		Scheme: "https",
		Host:   "server-cert:8080",
		Path:   "server",
	}
	log.Println(u.String())
	// req, err := http.NewRequest(http.MethodGet, "https://localhost:8080/server", nil)
	req, err := http.NewRequest(http.MethodGet, u.String(), nil)
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

 */
