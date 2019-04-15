package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Persion is type of persion
type Persion struct {
	ID        string   `json:"id,omitempty"`
	FirstName string   `json:"fistname,omitempty"`
	LastName  string   `json:"lastname,omitempty"`
	Address   *Address `json:"address,omitempty"`
}

// Address type of address
type Address struct {
	City  string `json:"city,omitempty"`
	State string `json:"state,opmitempty"`
}

var people []Persion

// Our main function
func main() {

	people = append(people, Persion{ID: "1", FirstName: "John", LastName: "Doe", Address: &Address{City: "city x", State: "america"}})
	people = append(people, Persion{ID: "2", FirstName: "Steve", LastName: "jon", Address: &Address{City: "city y", State: "india"}})
	people = append(people, Persion{ID: "3", FirstName: "what", LastName: "hell"})
	var randomKey = Key()
	print(randomKey)
	router := mux.NewRouter()
	router.HandleFunc("/people", GetPeople).Methods("GET")
	router.HandleFunc("/people/{id}", GetPersion).Methods("GET")
	router.HandleFunc("/people/{id}", AddPersion).Methods("POST")
	router.HandleFunc("/people/{id}", DeletePersion).Methods("DELETE")
	log.Fatal(http.ListenAndServe(":8000", router))

}

// GetPeople get people
func GetPeople(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(people)
}

// GetPersion get persion
func GetPersion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// AddPersion does add persion to list
func AddPersion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var persion Persion
	_ = json.NewDecoder(r.Body).Decode(&persion)
	persion.ID = params["id"]
	people = append(people, persion)
	json.NewEncoder(w).Encode(people)
}

// DeletePersion does delete persion
func DeletePersion(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id = params["id"]
	for index, item := range people {
		if item.ID == id {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(people)

}
