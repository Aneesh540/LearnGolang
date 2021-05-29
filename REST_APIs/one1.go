package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Author *Author `json:"author"`
}

type Author struct {
	First string `json:"firstname"`
	Last  string `json:"lastname"`
}

func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

func getBooksId(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	w.Header().Set("Content-Type", "application/json")

	for _, item := range books {
		if item.ID == params["id"] {
			w.WriteHeader(200)
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(404)
	json.NewEncoder(w).Encode(&Book{})
}

func addBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var b Book
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		w.WriteHeader(400)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	books = append(books, b)
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(&b)

}

var books []Book

func main() {
	books = append(books, Book{ID: "1", Title: "Noragami", Author: &Author{First: "Robert", Last: "Kiyo"}})
	books = append(books, Book{ID: "2", Title: "South park", Author: &Author{First: "rick", Last: "morty"}})
	books = append(books, Book{ID: "3", Title: "Family Guy", Author: &Author{First: "Lois", Last: "Griffin"}})

	// fmt.Println(books)

	r := mux.NewRouter()

	// handling routes
	r.HandleFunc("/books", getBook).Methods("GET")
	r.HandleFunc("/books/{id}", getBooksId).Methods("GET")
	r.HandleFunc("/books/insert", addBook).Methods("POST")
	log.Fatal(http.ListenAndServe(":1234", r))
}