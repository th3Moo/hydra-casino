package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    port := os.Getenv("PORT")
    if port == "" {
        port = "8082"
    }

    http.HandleFunc("/wallet", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "wallet ok")
    })

    log.Printf("Wallet service running on :%s", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
