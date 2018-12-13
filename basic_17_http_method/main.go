package main

import "net/http"
import "fmt"

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        if r.Method == "POST" {
            fmt.Fprintln(w, "Post")
        } else if r.Method == "GET" {
            fmt.Fprintln(w, "Get")
        } else {
            http.Error(w, "", http.StatusBadRequest)
        }
    })

    fmt.Println("server started at localhost:8080")
    http.ListenAndServe(":8080", nil)
}
