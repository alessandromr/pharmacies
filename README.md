# Pharmacies

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

Full test (Unit + Integration):
1. `docker build --tag jsonrpctest . -f test.dockerfile`
2. `docker run  --rm jsonrpctest:latest`

## Application Structure

The application structure is inspired by the `Standard Go Project Layout`.  

```
├── adapter
│   └── web
├── cmd
│   └── jsonrpc
├── config
├── datalayer
│   ├── memory
├── internal
│   ├── coordinates
│   └── rpcserver
├── model
├── presentation
│   └── jsonrpc
```

1. **adapter** contains the logic for external services comunication. An example is the logic to retrieve pharmacies list from the external source.
2. **cmd** contains all the entrypoint of the application divided by directory
3. **config** contains all the configuration files template or the default configs
4. **datalayer** contains all the logic to retrieve data from data sources. The actually data source is abstracted with an interface.
5. **internal** contains the unexported logic
6. **model** contains all the definitions of the entity involved
7. **presentation** containes all the logic and the abstraction to expose information on differents channels (RPC, REST, GRPC, ...)