package api

import (
	"log"
	"net/http"

	"github.com/dant89/go-microservice-api/src/api/internal/user"
	"github.com/julienschmidt/httprouter"
)

func Router() {
	router := httprouter.New()
	router.GET("/users", user.GetUsers)

	log.Fatal(http.ListenAndServe(":8080", router))
}
