package repository

import (
	"errors"

	"github.com/gocql/gocql"
)

type User struct {
	Id    string
	Name  string
	Email string
}

func GetUsers() []User {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "go_gopher_api"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()

	// TODO auto generate sample data if not exists
	//
	// CREATE KEYSPACE IF NOT EXISTS go_gopher_api
	//  WITH replication = {'class': 'SimpleStrategy',
	//                    'replication_factor' : 1}
	//
	// CREATE TABLE "go_gopher_api"."user" ("id" TIMEUUID, PRIMARY KEY (id));
	// ALTER TABLE "go_gopher_api"."user" ADD "name" TEXT;
	// ALTER TABLE "go_gopher_api"."user" ADD "email" TEXT;

	//var userCount int
	//if err := session.Query(`SELECT COUNT(id) FROM user`).Consistency(gocql.One).Scan(&userCount); err != nil {
	//	log.Fatal(err)
	//}

	// Populate data
	//if userCount == 0 {
	//	if err := session.Query(`INSERT INTO user (id, name, email) VALUES (?, ?, ?)`,
	//		gocql.TimeUUID(), "bob", "bob@bob.com").Exec(); err != nil {
	//		log.Fatal(err)
	//	}
	//}

	var id gocql.UUID
	var name string
	var email string

	var usersArray []User

	iter := session.Query(`SELECT id, name, email FROM user`).Iter()
	for iter.Scan(&id, &name, &email) {
		formattedUser := User{id.String(), name, email}
		usersArray = append(usersArray, formattedUser)
	}

	return usersArray
}

func GetUser(userId string) (User, error) {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "go_gopher_api"
	cluster.Consistency = gocql.Quorum
	session, _ := cluster.CreateSession()

	var id gocql.UUID
	var name string
	var email string

	userUUID, err := gocql.ParseUUID(userId)
	if err != nil {
		formattedErr := errors.New("User not found")
		return User{}, formattedErr
	}

	if err := session.Query(`SELECT id, name, email FROM user WHERE id = ?`, userUUID).Consistency(gocql.One).Scan(&id, &name, &email); err != nil {
		formattedErr := errors.New("User not found")
		return User{}, formattedErr
	}

	return User{id.String(), name, email}, nil
}
