package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/login", loginHandler).Methods("POST")
    r.HandleFunc("/health", healthHandler).Methods("GET")

    port := os.Getenv("PORT")
    if port == "" {
        port = "8081"
    }
    log.Printf("Auth service running on :%s", port)
    log.Fatal(http.ListenAndServe(":"+port, r))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "user": "demo",
    })
    tokenString, _ := token.SignedString(jwtSecret)
    fmt.Fprintf(w, tokenString)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "ok")
}
