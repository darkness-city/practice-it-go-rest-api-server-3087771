package helloworld

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	"github.com/gorilla/mux"
)

type App struct {
	DB 		*sql.DB
	Port 	string
}

func getRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a GET\n")
}

func postRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a POST\n")
}

func deleteRequest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is a Delete\n")
}

func (a *App) Initialize() {
	var err error
	a.DB , err = sql.Open("sqlite3", "../practiceit.db")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func (a *App) allProducts(w http.ResponseWriter, r *http.Request) {
	products, err := getProducts(a.DB)
	if err != nil {
		fmt.Printf("getProducts error: %s\n",err.Error())
	}
}
func (a *App) Run() {
	r := mux.NewRouter()
	r.HandleFunc("/products", a.allProducts).Methods("GET")
	r.HandleFunc("/product/{id}", a.fetchProduct).Methods("GET")
	r.HandleFunc("/", getRequest ).Methods("GET")
	r.HandleFunc("/", postRequest).Methods("POST")
	r.HandleFunc("/", deleteRequest).Methods("DELETE")

	http.Handle("/", r)
	fmt.Println("Server Started and listening on localhost",a.Port)
	log.Fatal(http.ListenAndServe(a.Port, nil))
	// http.HandleFunc("/", handle)
	// fmt.Println("Server started and listening on localhost",a.Port)
	// log.Fatal(http.ListenAndServe(a.Port, nil))
}