package main

import (
	"flag"
	"log"
	"net/http"

	api "github.com/dant89/go-gopher-api/src/api/router"
)

func main() {
	portParam := flag.String("p", "8080", "The localhost port to run the API web server.")
	router := api.Router()

	log.Fatal(http.ListenAndServe(":"+*portParam, router))
}
