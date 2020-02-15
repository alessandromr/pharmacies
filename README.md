# Test Closest Pharmacies

## How to start the application

### Requirements
* Docker

### Steps

1. Build the container `docker build --tag jsonrpc`
2. Start the container `docker run --rm -p 8080:8080 --name jsonrpc jsonrpc:latest`

## Testing
Test execution can be done on your local machine or in a Docker container (without installing anything beside of Docker).  

### Testing on local machine

* Unit Test `make test`
* Integration Test `make integration`

Test coverage 
```bash
go test --cover ./...
```

### Testing with docker

Complete test:
1. `docker build --tag jsonrpctest . -f test.dockerfile`
2. `docker run  --rm jsonrpctest:latest`

## Application Structure

```
├── adapter
│   └── web
├── datalayer
│   └── memory
├── model
├── presentation
```

1. adapter directory contains the logic for external services comunication. An example is the logic to retrieve pharmacies list from the external source.
2. datalayer directory contains all the logic to retrieve data from data sources. The actually data source is abstracted with an interface.
3. model directory contains all the definitions of the entity involved
3. presentation containes all the logic and the abstraction to expose information on differents channels