package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ToDo struct {
	ID   int    `json:"id"`
	Task string `json:"task"`
}

var toDoList []ToDo
var numberOfToDos int

func addToDo(w http.ResponseWriter, r *http.Request) {
	// Function to add a new todo item
	var myToDo ToDo
	// Parse the request body to get the new todo item
	// Add the new todo item to the list
	err := json.NewDecoder(r.Body).Decode(&myToDo)
	if err != nil {
		http.Error(w, "Invalid Input Request", http.StatusBadRequest)
		return
	}

	//After decoding the request body
	// Now we can set the ID of the new todo item

	myToDo.ID = numberOfToDos
	numberOfToDos += 1

	// Append the new todo item to the list
	toDoList = append(toDoList, myToDo)

	// Return the new todo item as a JSON response
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(myToDo)
}

// Handler function to see the list of todo items
func getToDos(w http.ResponseWriter, r *http.Request) {
	// Function to get all todo items
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(toDoList)
}

func main() {
	http.HandleFunc("/todos", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			addToDo(w, r)
		} else if r.Method == http.MethodGet {
			getToDos(w, r)
		}
	})

	fmt.Println("Server is running at 8080 port")
	http.ListenAndServe(":8080", nil)
}
