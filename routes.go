package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	var books []Book
	for _, v := range bookStore {
		books = append(books, v)
	}
	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(books)

	if err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func PostBookHandler(w http.ResponseWriter, r *http.Request) {
	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)

	if err != nil {
		panic(err)
	}

	id++
	k := strconv.Itoa(id)
	bookStore[k] = book

	w.Header().Set("Content-Type", "application/json")
	j, err := json.Marshal(book)
	w.WriteHeader(http.StatusCreated)
	w.Write(j)
}

func PutBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var bookUpdate Book
	err := json.NewDecoder(r.Body).Decode(&bookUpdate)

	if err != nil {
		panic(err)
	}

	if _, ok := bookStore[id]; ok {
		delete(bookStore, id)
		bookStore[id] = bookUpdate
	} else {
		log.Printf("Not found ID: %s", id)
	}

	w.WriteHeader(http.StatusNoContent)
}

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if _, ok := bookStore[id]; ok {
		delete(bookStore, id)
	} else {
		log.Printf("Not found ID: %s", id)
	}

	w.WriteHeader(http.StatusNoContent)
}
