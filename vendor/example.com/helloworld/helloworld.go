package helloworld

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

type App struct {
	DB 		*sql.DB
	Port 	string
}

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World\n")
}

func (a *App) Initialize() {
	var err error
	a.DB , err = sql.Open("sqlite3", "../practiceit.db")
	if err != nil {
		log.Fatal(err.Error())
	}
}
func (a *App) Run() {
	http.HandleFunc("/", handle)
	fmt.Println("Server started and listening on localhost",a.Port)
	log.Fatal(http.ListenAndServe(a.Port, nil))
}