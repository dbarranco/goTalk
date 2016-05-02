package main

import (
    "net/http" 
    "fmt"
)

func handler(w http.ResponseWriter, r *http.Request) { 
    fmt.Println("Manejador HTTP")
    fmt.Fprintf(w, "¡Hola mundo!")
}

func main() {
    http.HandleFunc("/", handler) 
    http.ListenAndServe("localhost:9999", nil) 
}