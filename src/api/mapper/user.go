package mapper

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
)

const userCsv string = "data-source/users.csv"

type User struct {
	Id    string
	Name  string
	Email string
}

func FetchUsers() []User {

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

	return formattedUsers
}

func FetchUser(id string) (User, error) {
	users := FetchUsers()

	var formattedUser User
	for _, user := range users {
		if id == user.Id {
			formattedUser = user
			break
		}
	}

	var err error
	fmt.Println("user", formattedUser.Id)
	if formattedUser.Id == "" {
		err = errors.New("User not found")
	}

	return formattedUser, err
}
