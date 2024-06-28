package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Starting library api ...")
	r := mux.NewRouter()

	r.HandleFunc("/books", GetBooks).Methods("GET")
	r.HandleFunc("/books/{id}", GetBookById).Methods("GET")
	r.HandleFunc("/books", CreateOneBook).Methods("POST")
	r.HandleFunc("/books/{id}", UpdateBookById).Methods("PUT")
	r.HandleFunc("/books/{id}", DeleteBookById).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", r))
}
