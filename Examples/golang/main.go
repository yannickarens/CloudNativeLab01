package main

import (
"fmt"
"net/http"
"log"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
fmt.Fprintf(w, "<html><head></head><body bgcolor='red'></body></html>")
}

func main() {
http.HandleFunc("/", sayhello) // set router
err := http.ListenAndServe(":8080", nil) // set listen port
if err != nil {
log.Fatal("ListenAndServe: ", err)
}
}