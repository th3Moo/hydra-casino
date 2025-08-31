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
        port = "8083"
    }

    http.HandleFunc("/tx", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "tx ok")
    })

    log.Printf("Transaction service running on :%s", port)
    log.Fatal(http.ListenAndServe(":"+port, nil))
}
