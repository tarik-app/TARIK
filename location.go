package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "text/html")
    if r.URL.Path == "/" {
      fmt.Fprintf(w, "<h1>Hi there, Curious where you are, me too!</h1>")
      fmt.Fprintf(w, "You are at %s!\n", r.URL.Path[1:])
      fmt.Fprintf(w, "Let's see what you know about %s", r.URL.Path[1:])

    } else if r.URL.Path == "/contact" {
      fmt.Fprintf(w, "To get in touch, please send an eamil to <a href=\"mailto:support@lenslocked.com\"> support@lenslocked.com</a>.")
    } else {
      w.WriteHeader(http.StatusNotFound)
      fmt.Fprint(w, "<h1>Invalid Page</h1>")
    }



}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
