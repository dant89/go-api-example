package api

import (
	"github.com/dant89/go-gopher-api/src/api/router/user"
	"github.com/julienschmidt/httprouter"
)

func Router() *httprouter.Router {
	router := httprouter.New()
	router.GET("/users", user.GetUsers)
	router.GET("/users/:id", user.GetUser)
	return router
}
