package main

import (
	"flag"

	api "github.com/dant89/go-microservice-api/src/api/router"
)

func main() {
	portParam := flag.String("p", "8080", "The localhost port to run the API web server.")
	api.Router(*portParam)
}
