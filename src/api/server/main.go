package main

import (
	"flag"
	"log"
	"net/http"

	router "github.com/dant89/go-gopher-api/src/api/router"
)

func main() {
	portParam := flag.String("p", "8080", "The localhost port to run the API web server.")
	router := router.Router()

	log.Fatal(http.ListenAndServe(":"+*portParam, router))
}
