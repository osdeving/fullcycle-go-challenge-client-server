package server

import (
	"fmt"
	"net/http"
)

func Start() {
    fmt.Println("Server is running...")
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, "Hello, World!")
    })
    http.ListenAndServe(":8080", nil)
}
