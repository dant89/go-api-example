package user

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dant89/go-microservice-api/src/api/mapper"
	"github.com/julienschmidt/httprouter"
)

type ApiError struct {
	Status int
	Error  string
}

func GetUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	users := mapper.FetchUsers()

	json, err := json.Marshal(users)
	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, string(json))
}

func GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userId := ps.ByName("id")
	user, err := mapper.FetchUser(userId)
	if err != nil {
		errorResponse(w, 404, "Specified user not found.")
		return
	}

	json, err := json.Marshal(user)
	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, string(json))
}

func errorResponse(w http.ResponseWriter, errorCode int, errorMsg string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(errorCode)

	error := ApiError{errorCode, errorMsg}

	json.
		NewEncoder(w).
		Encode(error)
}
