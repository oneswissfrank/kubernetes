package main

import (
    "fmt"
    "log"
    "io/ioutil"
    "net/http"
    "time"
    "os"
    "strconv"
)

func homePage(w http.ResponseWriter, r *http.Request){
    response, err := http.Get("http://backend.backend:8080/hello")
    if err != nil {
        fmt.Print(err.Error())
        os.Exit(1)
    }
    responseData, err := ioutil.ReadAll(response.Body)
    if err != nil {
        log.Fatal(err)
    }
    now := time.Now()
    seconds := now.Unix()
    w.Header().Set("My-Custom-Header", strconv.FormatInt(seconds, 10))
    fmt.Fprintf(w, string(responseData))
}

func handleRequests() {
    http.HandleFunc("/", homePage)
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
    handleRequests()
}
