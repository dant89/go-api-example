module github.com/dant89/go-gopher-api

go 1.16

require (
	github.com/go-redis/redis/v8 v8.8.2
	github.com/gocql/gocql v0.0.0-20210425135552-909f2a77f46e
	github.com/joho/godotenv v1.3.0
	github.com/julienschmidt/httprouter v1.3.0
)

replace github.com/dant89/go-microservice-api v0.0.0-20210501171054-5d7fdacd9da9 => github.com/dant89/go-gopher-api v0.0.0-20210501171054-5d7fdacd9da9
