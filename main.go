package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Status: Active")
	fmt.Fprintln(w, "User: Priyanshu")
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the DevOps Bootcamp!")
}

func main() {
	// 1. .env file load karo
	godotenv.Load()

	// 2. PORT variable uthao
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Agar .env nahi mili to backup
	}

	http.HandleFunc("/", rootHandler)
	http.HandleFunc("/health", healthHandler)

	log.Printf("Server starting on port %s...", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
