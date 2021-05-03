package router

import (
	"github.com/julienschmidt/httprouter"
)

func Router() *httprouter.Router {
	router := httprouter.New()
	router.GET("/users", GetUsers)
	router.GET("/users/:id", GetUser)
	return router
}
