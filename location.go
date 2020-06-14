package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type TarikPage struct {
	HistoryOf   string
	Description string
}

func TarikHandler(w http.ResponseWriter, r *http.Request) {
	tarik := TarikPage{HistoryOf: "ChinaTown", Description: "For the first Chinatown in the world...."}
	t, _ := template.ParseFiles("index.html")
	t.Execute(w, tarik)
}

func login(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		fmt.Println("method in GET:", r.Method) //get request method
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		r.ParseForm()
		// logic part of log in
		fmt.Println("method in POST:", r.Method) //get request method
		fmt.Println("Location-Query:", r.Form["location"])
	}
}

func main() {
	r := mux.NewRouter()
	// r.HandleFunc("/tarik", TarikHandler)
	r.HandleFunc("/tourist-site", TouristAttractionHandler)
	// r.HandleFunc("/login", login)
	log.Fatal(http.ListenAndServe(":8100", r))
}
