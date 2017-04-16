package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Item struct definition with its json field name
type Item struct {
	ID   string `json:"id, omitempty"`
	Name string `json:"name, omitempty"`
}

// ErrorMessage struct definition
type ErrorMessage struct {
	Error string `json:"error, omitempty"`
}

var items []Item

func getItemsEndpoint(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(items)
}

func getItemEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range items {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	errorHandler(w, r, 404, "Item not found.")
}

func createItemEndpoint(w http.ResponseWriter, r *http.Request) {
	var item Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		errorHandler(w, r, 400, "Invalid request body.")
		return
	}
	items = append(items, item)
	json.NewEncoder(w).Encode(items)
}

func deleteItemEndpoint(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range items {
		if item.ID == params["id"] {
			// replace items with items without current ID
			items = append(items[:index], items[index+1:]...)
			json.NewEncoder(w).Encode(items)
			return
		}
	}
	errorHandler(w, r, 404, "Item not found.")
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int, message string) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(ErrorMessage{Error: message})
}

func main() {
	// init sample item data
	items = append(items, Item{ID: "1", Name: "Sample item name"})

	// routing
	router := mux.NewRouter()
	router.HandleFunc("/items", getItemsEndpoint).Methods("GET")
	router.HandleFunc("/items", createItemEndpoint).Methods("POST")
	router.HandleFunc("/items/{id}", getItemEndpoint).Methods("GET")
	router.HandleFunc("/items/{id}", deleteItemEndpoint).Methods("DELETE")

	// start server
	log.Println("Starting...")
	log.Fatal(http.ListenAndServe(":3000", router))
}
