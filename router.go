package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func getRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a GET")
}

func postRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a POST")
}

func deleteRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a Delete")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getRequest ).Methods("GET")
	r.HandleFunc("/", postRequest).Methods("POST")
	r.HandleFunc("/", deleteRequest).Methods("DELETE")

	http.Handle("/", r)
	fmt.Println("Server Started and listening on localhost:9003")
	log.Fatal(http.ListenAndServe(":9003", nil))
}