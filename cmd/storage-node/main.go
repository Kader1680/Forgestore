package main

import (
    "log"
    "net/http"
)

func main() {
    mux := http.NewServeMux()

    mux.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "application/json")
        w.Write([]byte(`{"service":"storage-node","status":"healthy"}`))
    })

    log.Println("Storage Node running on :8081")

    if err := http.ListenAndServe(":8081", mux); err != nil {
        log.Fatal(err)
    }
}