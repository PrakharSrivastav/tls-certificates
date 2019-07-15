package main

import (
	"fmt"
	"log"
	"net/http"
)

// Handle
// Handler
// HandleFunc
// ServeMux
func main() {
	fmt.Println("starting server")
	routes := new(router)

	mux := http.NewServeMux()

	mux.HandleFunc("/", routes.routeRoot)
	mux.HandleFunc("/foo", routes.routeFoo)

	log.Fatal(http.ListenAndServe(":8080", mux))
}

type router struct{}

func (*router) routeRoot(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "root")
}

func (*router) routeFoo(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintf(w, "foo")
}
