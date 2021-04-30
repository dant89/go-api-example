# go-microservices-api

An example of how microservices can work in Golang using self contained packages.

The API is a very basic implementation that reads user data from a static csv file and formats it into json output.

## Usage
1. Start the API web server by running: `go run src/api/main.go`
2. Run an example query via the client: `go run src/client/example/request.go`
