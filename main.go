package main

import (
    "fmt"
    "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the DevOps Bootcamp!")
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Status: Active\nUser: Priyanshu")
}

func main() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/health", healthCheck)

    fmt.Println("Server is starting on port 8080...")
    http.ListenAndServe("0.0.0.0:8080", nil)
}
