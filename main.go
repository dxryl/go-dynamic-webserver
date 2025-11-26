package main

import (
    "fmt"
    "net/http"
    "time"
)

func handler(w http.ResponseWriter, r *http.Request) {
    currentTime := time.Now().Format("3:04 PM Monday Jan 2 2006")
    message := fmt.Sprintf("Hello! The current server time is: %s", currentTime)
    fmt.Fprintf(w, message)
}

func main() {
    http.HandleFunc("/", handler)

    fmt.Println("Starting server on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
