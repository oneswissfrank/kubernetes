package main

import (
    "fmt"
    "log"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
    id := r.Header.Get("My-Custom-Header")
    w.Header().Set("My-Custom-Header", id)
    fmt.Fprintf(w, "Hello from the backend!")
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
    handleRequests()
}
