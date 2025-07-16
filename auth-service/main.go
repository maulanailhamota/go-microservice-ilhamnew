package main

import (
    "fmt"
    "log"
    "os"
    "net/http"
    "github.com/joho/godotenv"
)

func main() {
    // Load .env
    err := godotenv.Load()
    if err != nil {
        log.Println("Tidak bisa memuat file .env, pakai default environment.")
    }

    ownerName := os.Getenv("OWNER_NAME")
    if ownerName == "" {
        ownerName = "Unknown Owner"
    }

    port := os.Getenv("PORT")
    if port == "" {
        port = "8081"
    }

    fmt.Printf("%s (%s) is running on port %s...\n", os.Getenv("SERVICE_NAME"), ownerName, port)

    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "%s Service - Owner: %s", os.Getenv("SERVICE_NAME"), ownerName)
    })

    log.Fatal(http.ListenAndServe(":" + port, nil))
}
