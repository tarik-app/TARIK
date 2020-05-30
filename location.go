package main

import (
    "fmt"
    "log"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "<h1>Hi there, Curious where you are, me too!</h1>")
    fmt.Fprintf(w, "You are at %s!\n", r.URL.Path[1:])
    fmt.Fprintf(w, "Let's see what you know about %s", r.URL.Path[1:])
}

func main() {
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
