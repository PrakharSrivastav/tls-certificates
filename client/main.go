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

func main() {

	c := getClient()
	r := getRequest()

	res, err := c.Do(r)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println("error reading body")
		log.Fatal(err)
	}
	log.Printf("%#v\n", string(body))
}

func getRequest() *http.Request {
	r := http.Request{
		URL:    getURL(),
		Method: http.MethodGet,
		Body:   nil,
	}
	return &r
}

func getClient() *http.Client {
	// root ca
	rootPem, err := ioutil.ReadFile("../minica.pem")
	if err != nil {
		log.Fatal(err)
	}
	roots := x509.NewCertPool()
	roots.AppendCertsFromPEM(rootPem)

	// key pair
	pair, err := tls.LoadX509KeyPair("cert.pem", "key.pem")
	if err != nil {
		log.Fatal(err)
	}
	c := http.Client{
		Transport: &http.Transport{
			TLSHandshakeTimeout: time.Second * 30,
			TLSClientConfig: &tls.Config{
				RootCAs:      roots,
				Certificates: []tls.Certificate{pair},
			},
		},
		Timeout: time.Second * 30,
	}

	return &c
}

func getURL() *url.URL {
	u := url.URL{
		Host:   "server-cert:8080",
		Path:   "",
		Scheme: "https",
	}

	return &u
}
