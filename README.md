# go-gopher-api

A lightweight Golang API framework. Aims to be easily extendable and customisable.


## Usage
1. Start the Docker containers: `docker-compose up -d`
2. Create a default Cassandra keyspace and table (this will be automated in the future):
```
CREATE KEYSPACE IF NOT EXISTS go_gopher_api
    WITH replication = {'class': 'SimpleStrategy',
                        'replication_factor' : 1}

CREATE TABLE "go_gopher_api"."user" ("id" TIMEUUID, PRIMARY KEY (id));
ALTER TABLE "go_gopher_api"."user" ADD "name" TEXT;
ALTER TABLE "go_gopher_api"."user" ADD "email" TEXT;
```
3. Start the API web server: `go run src/server/main.go`
4. Run the example queries via the client: `go run src/client/example/request.go`

## External libraries
This project is making use of the following third party libraries, check them out!
- [Go Redis client](https://github.com/go-redis/redis/v8)
- [Go Cassandra client (gocql)](https://github.com/gocql/gocql)
- [Go HTTP Router](https://github.com/julienschmidt/httprouter)
- [Go DotEnv parser](https://github.com/joho/godotenv)

## Why does this exist?
I'm building this to improve my Go programming. If people find it useful, that's a big bonus!
