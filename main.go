package main

import (
    "fmt"
    "log"
    "net/http"
)

var version = "v0"

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Welcome! Version %s", version)
    })
    fmt.Printf("App starts, version: %s \n", version)
    log.Fatal(http.ListenAndServe(":8500", nil))
} 
