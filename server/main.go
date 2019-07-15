package main

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "This is sparta")
	})

	caCert, err := ioutil.ReadFile("../minica.pem")
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	c := tls.Config{
		ClientCAs:  caCertPool,
		ClientAuth: tls.RequireAndVerifyClientCert,
	}
	c.BuildNameToCertificate()

	s := &http.Server{
		Addr:      ":8080",
		TLSConfig: &c,
		Handler:   mux,
	}

	log.Fatal(s.ListenAndServeTLS("cert.pem", "key.pem"))
}
