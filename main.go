package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Book struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      int    `json:"author"`
	Year        int    `json:"year"`
}

var bookStore = make(map[string]Book)
var id int

func main() {
	r := mux.NewRouter().StrictSlash(false)

	r.HandleFunc("/v1/api/books", GetBooksHandler).Methods("GET")
	r.HandleFunc("/v1/api/books", PostBookHandler).Methods("POST")
	r.HandleFunc("/v1/api/books/{id}", PutBookHandler).Methods("PUT")
	r.HandleFunc("/v1/api/books/{id}", DeleteBookHandler).Methods("DELETE")

	srv := &http.Server{
		Addr:           ":8000",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	fmt.Println("Listing on port 8000")
	log.Fatal(srv.ListenAndServe())
}
