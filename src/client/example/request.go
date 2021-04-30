package main

import (
	"fmt"

	"github.com/dant89/go-microservice-api/src/client"
)

func main() {

	host := "http://127.0.0.1:8080"

	myClient := client.NewClient(host)
	users := myClient.GetUsers()

	fmt.Println(users)
}
