package main

import (
    "log"
    "net/http"
)

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("GET /health", func(writer http.ResponseWriter, request *http.Request) {
        writer.Header().Set("Content-Type", "application/json")
        writer.Write([]byte(`{"service":"coordinator","status":"healthy"}`))
    })

    log.Println("Coordinator running on :8080")

    if err := http.ListenAndServe(":8080", mux); err != nil {
        log.Fatal(err)
    }
}