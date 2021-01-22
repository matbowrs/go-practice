package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Structs / Model
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Country   string `json:"country"`
}

// Init books var - slice (variable length array)
var books []Book

// Get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {

}
func createBook(w http.ResponseWriter, r *http.Request) {

}
func updateBook(w http.ResponseWriter, r *http.Request) {

}
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {
	// Initialize new router
	r := mux.NewRouter()

	// Mock data
	books = append(books, Book{ID: "1", Isbn: "1123434", Title: "Евгений Онегин", Author: &Author{Firstname: "Alexander", Lastname: "Pushkin", Country: "Russia"}})
	books = append(books, Book{ID: "2", Isbn: "2437865", Title: "Мастер и Маргерита", Author: &Author{Firstname: "Mikhail", Lastname: "Bulgakov", Country: "Russia"}})
	books = append(books, Book{ID: "3", Isbn: "3543123", Title: "La Nausée", Author: &Author{Firstname: "Jean-Paul", Lastname: "Sartre", Country: "France"}})

	// Create route handlers. This establishes endpoints for api
	// .Methods() used for what kind of request
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books/", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Run server
	// Wrapped in log.Fatal() for errors
	log.Fatal(http.ListenAndServe(":3000", r))
}
