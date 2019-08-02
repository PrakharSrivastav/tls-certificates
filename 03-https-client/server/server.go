package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	CertPath string = "../../00-certificates/server/cert.pem"
	KeyPath  string = "../../00-certificates/server/key.pem"
)

func main() {
	mux := http.NewServeMux()

	// add an endpoint
	mux.HandleFunc("/server", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "i am protected")
	})
	log.Println("starting server")
	log.Fatal(http.ListenAndServeTLS(":8080", CertPath, KeyPath, mux))
}
