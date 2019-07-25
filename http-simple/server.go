package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	// add an endpoint
	mux.HandleFunc("/server", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "i am naked")
	})

	log.Fatal(http.ListenAndServe(":8080", mux))
}
