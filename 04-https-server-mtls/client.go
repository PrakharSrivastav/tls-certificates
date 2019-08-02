package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

const (
	RootCertificatePath string = "../00-certificates/minica.pem"
	ClientCertPath      string = "../00-certificates/client/cert.pem"
	ClientKeyPath       string = "../00-certificates/client/key.pem"
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
				GetClientCertificate: func(info *tls.CertificateRequestInfo) (certificate *tls.Certificate, e error) {
					log.Println("request from server")
					c, err := tls.LoadX509KeyPair(ClientCertPath, ClientKeyPath)
					if err != nil {
						fmt.Printf("Error loading key pair: %v\n", err)
						return nil, err
					}
					return &c, nil
				},
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
