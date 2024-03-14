package main

import (
	"fmt"
	"net/http"
)

func main() {
	// register a request handler
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	// listen on a port for http connections
	http.ListenAndServe(":8080", nil)

}
