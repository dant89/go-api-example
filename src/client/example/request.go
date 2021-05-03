package main

import (
	"fmt"

	"github.com/dant89/go-gopher-api/src/client"
)

func main() {

	host := "http://127.0.0.1:8080"

	myClient := client.NewClient(host)

	fmt.Print("\nAn example of all users output\n")
	users := myClient.GetUsers()
	fmt.Print(users, "\n\n")

	fmt.Print("An example of a specific user output\n")
	user := myClient.GetUser("939d5194-a051-11eb-bcbc-0242ac130002")
	fmt.Print(user, "\n\n")

	fmt.Print("An example of a specific not found user output\n")
	notFoundUser := myClient.GetUser("zzzzz")
	fmt.Print(notFoundUser, "\n")
}
