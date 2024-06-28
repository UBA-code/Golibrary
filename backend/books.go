package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

var books []Book = []Book{
	{
		Name:     "new world",
		Author:   "UBA",
		Price:    999,
		Category: "Theory",
	},
	{
		Name:     "Art of typescript",
		Author:   "UBA",
		Price:    10,
		Category: "Programing",
	},
}

type Book struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Author   string `json:"author"`
	Price    int    `json:"price"`
	Category string `json:"category"`
}

func (b *Book) IsEmpty() bool {
	return b.Name == "" || b.Author == "" || b.Category == ""
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get all books")
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(books)
}

func GetBookById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("get book by id")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for _, book := range books {
		if book.Id == params["id"] {
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode("no book found with the given id: " + params["id"])
}

func UpdateBookById(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update existing book")
	w.Header().Set("Content-Type", "application/json")
	if r.Body == nil {
		json.NewEncoder(w).Encode("no data in body")
		return
	}
	params := mux.Vars(r)
	var newBook Book
	err := json.NewDecoder(r.Body).Decode(&newBook)
	if err != nil {
		json.NewEncoder(w).Encode("Error while decode object")
	}
	if newBook.IsEmpty() {
		json.NewEncoder(w).Encode("Please give a valid object")
	}
	for index, book := range books {
		if book.Id == params["id"] {
			books = append(books[:index], books[index+1:]...)
			books = append(books, newBook)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}
