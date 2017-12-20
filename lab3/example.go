package main

import (
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
    fmt.Fprintf(w,"%s",r.Method)
}

func main() {
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8084", nil)
}
