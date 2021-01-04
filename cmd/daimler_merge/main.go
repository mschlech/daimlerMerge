package main

import (
	"log"
	"net/http"
)

func main() {
	log.Printf("start service")
	router := NewRouter()
	log.Fatal(http.ListenAndServe(":3000", router))
	//prometheus should only be visible from inside the infra will not work within this Dockerfile approach! see routes for prometheus metrics

}
