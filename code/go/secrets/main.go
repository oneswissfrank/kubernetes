package main

import (
    "fmt"
    "os"
    "log"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
    envVarKey := os.Getenv("ENV_VAR_KEY")
    fmt.Fprintf(w, "The %s value is: %s", envVarKey, os.Getenv(envVarKey))
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
    handleRequests()
}
