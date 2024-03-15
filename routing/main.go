package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/books/{title}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		title := vars["title"] // book title slug
		page := vars["page"]   // the page

		fmt.Fprintf(w, "You've requested the book: %s on page %s\n", title, page)
	})

	r.HandleFunc("/books/{title}", CreateBook).Methods("POST") // restricts request handler to specific HTTP methods
	r.HandleFunc("/books/{title}", ReadBook).Methods("GET")
	r.HandleFunc("/books/{title}", UpdateBook).Methods("PUT")
	r.HandleFunc("/books/{title}", DeleteBook).Methods("DELETE")

	r.HandleFunc("/books/{title}", BookHandler).Host("www.mybookstore.com") // restrict the request handler to specific hostnames or subdomains

	bookrouter := r.PathPrefix("/books").Subrouter() // resticts the request handler to specific path prefixes
	bookrouter.HandleFunc("/", AllBooks)
	bookrouter.HandleFunc("/{title}", GetBook)

	http.ListenAndServe(":8080", r)

}
