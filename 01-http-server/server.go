package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// mux and the endpoint
	mux := http.NewServeMux()
	mux.HandleFunc("/server", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Protect Me...")
	})

	// start the http server
	log.Println("starting server")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
