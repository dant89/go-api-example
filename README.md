# go-microservices-api

An example of how microservices can work in Golang using self contained packages.

The API is a very basic implementation that reads user data from a static csv file and formats it into json output.

## Usage
1. Start the Docker Cassandra container: `docker-compose up`
2. Create a default Cassandra keyspace and table (this will be automated in the future):
```
	CREATE KEYSPACE IF NOT EXISTS go_microservice_api
	    WITH replication = {'class': 'SimpleStrategy',
	                       'replication_factor' : 1}

	CREATE TABLE "go_microservice_api"."user" ("id" TIMEUUID, PRIMARY KEY (id));
	ALTER TABLE "go_microservice_api"."user" ADD "name" TEXT;
	ALTER TABLE "go_microservice_api"."user" ADD "email" TEXT;
```
3. Start the API web server: `go run src/api/main.go`
4. Run the example queries via the client: `go run src/client/example/request.go`
