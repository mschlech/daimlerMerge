package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("start service")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":3000", router))
}
