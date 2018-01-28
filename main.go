package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book Struct (MODEL) - Like a class
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
type Author struct {
	Firtname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

// Init books var as a slice(length array) Book struct
var books []Book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBookById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Get params
	// Loop through books and find by id

	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	json.NewEncoder(w).Encode(&Book{})
}

func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book

	_ = json.NewDecoder(r.Body).Decode(&book)

	book.ID = strconv.Itoa(rand.Intn(1000000000))

	books = append(books, book)

	json.NewEncoder(w).Encode(book)

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}

func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(books)
}

func main() {

	books = append(books, Book{ID: "1", Isbn: "102831024182", Title: "Book 1", Author: &Author{Firtname: "Victor1", Lastname: "Costa"}})
	books = append(books, Book{ID: "2", Isbn: "109751912731", Title: "Book 2", Author: &Author{Firtname: "Victor2", Lastname: "Costa"}})
	books = append(books, Book{ID: "3", Isbn: "049759812398", Title: "Book 3", Author: &Author{Firtname: "Victor3", Lastname: "Costa"}})
	books = append(books, Book{ID: "4", Isbn: "198724969929", Title: "Book 4", Author: &Author{Firtname: "Victor4", Lastname: "Costa"}})
	books = append(books, Book{ID: "5", Isbn: "785447881189", Title: "Book 5", Author: &Author{Firtname: "Victor5", Lastname: "Costa"}})

	// Init mux router
	router := mux.NewRouter() // := means type inference

	// Route handler / Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBookById).Methods("GET")
	router.HandleFunc("/api/books", createBook).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))

}
