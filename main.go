package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Structs / Model
// Book struct
type Book struct {
	Id     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

// Author struct
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
	// Outputs json data
	json.NewEncoder(w).Encode(books)
}
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// Gets parameters
	params := mux.Vars(r)

	for _, item := range books {
		if item.Id == params["id"] {
			fmt.Println(item)
			json.NewEncoder(w).Encode(item)
			return
		}
	}

	// Outputs empty Book, since id was not found
	json.NewEncoder(w).Encode(&Book{})
}
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)

	// wrapped in strconv because this has to be a string
	book.Id = strconv.Itoa(rand.Intn(10000000))

	// append to our slice
	books = append(books, book)

	// output new book, json data
	json.NewEncoder(w).Encode(book)
}
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, item := range books {
		if item.Id == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)

			// wrapped in strconv because this has to be a string
			book.Id = strconv.Itoa(rand.Intn(10000000))

			// append to our slice
			books = append(books, book)
			json.NewEncoder(w).Encode(books)

			return
		}
	}

	json.NewEncoder(w).Encode(books)
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.Id == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}

	json.NewEncoder(w).Encode(books)
}

func main() {
	// Initialize new router
	r := mux.NewRouter()

	// Mock data
	books = append(books, Book{Id: "1", Isbn: "1123434", Title: "Евгений Онегин", Author: &Author{Firstname: "Alexander", Lastname: "Pushkin", Country: "Russia"}})
	books = append(books, Book{Id: "2", Isbn: "2437865", Title: "Мастер и Маргерита", Author: &Author{Firstname: "Mikhail", Lastname: "Bulgakov", Country: "Russia"}})
	books = append(books, Book{Id: "3", Isbn: "3543123", Title: "La Nausée", Author: &Author{Firstname: "Jean-Paul", Lastname: "Sartre", Country: "France"}})

	// Create route handlers. This establishes endpoints for api
	// .Methods() used for what kind of request
	r.HandleFunc("/api/books", getBooks).Methods("GET")
	r.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/api/books", createBook).Methods("POST")
	r.HandleFunc("/api/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/api/books/{id}", deleteBook).Methods("DELETE")

	// Run server
	// Wrapped in log.Fatal() for errors
	log.Fatal(http.ListenAndServe(":3000", r))
}
