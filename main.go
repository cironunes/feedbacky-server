package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"
)

func main() {
	router := NewRouter()
	handler := cors.Default().Handler(router)

	log.Fatal(http.ListenAndServe(":4222", handler))
}
