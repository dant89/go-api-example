package user

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/julienschmidt/httprouter"
)

const userCsv string = "data-source/users.csv"

type User struct {
	Id    string
	Name  string
	Email string
}

func GetUsers(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	f, err := os.Open(userCsv)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	users, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	var formattedUsers []User
	for _, user := range users {
		formattedUser := User{user[0], user[1], user[2]}
		formattedUsers = append(formattedUsers, formattedUser)
	}

	json, err := json.Marshal(formattedUsers)
	if err != nil {
		panic(err)
	}
	fmt.Fprint(w, string(json))
}
