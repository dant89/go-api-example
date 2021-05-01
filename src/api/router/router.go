package api

import (
	"log"
	"net/http"

	"github.com/dant89/go-microservice-api/src/api/router/user"
	"github.com/julienschmidt/httprouter"
)

func Router(port string) {
	router := httprouter.New()
	router.GET("/users", user.GetUsers)
	router.GET("/users/:id", user.GetUser)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
