package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func home(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  fmt.Fprint(w, "<h1>Welcome to my awseome site!</h1>")
}

func contact(w http.ResponseWriter, r *http.Request) {
  w.Header().Set("Content-Type", "text/html")
  fmt.Fprintf(w, "To get in touch, please send an eamil to <a href=\"mailto:support@lenslocked.com\"> support@lenslocked.com</a>.")
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/", home)
    r.HandleFunc("/contact", contact)
    log.Fatal(http.ListenAndServe(":8080", r))
}
