package main

import (
	"log"
	"net/http"
	)

func main() {
	log.Printf("start service")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
