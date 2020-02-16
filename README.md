# Pharmacies

The application retrieve data about pharmacies in Campania (Italy) on startup from an exposed list provided by Regione Campania.  
Data about pharmacies is refreshed every 24 hours and persisted in memory.  

## How to start the application

### Requirements
* Docker

### Steps

1. Build the container `docker build --tag jsonrpc .`
2. Start the container `docker run --rm -p 8080:8080 --name jsonrpc jsonrpc:latest`


Application will be started on port `8080` inside the docker container. To change the port binding on your host simply change the first port in the docker run command. 
For example to use port `9901` you can use this command:
```docker
    docker run --rm -p 9901:8080 --name jsonrpc jsonrpc:latest
```

## Automated Testing
Test execution can be done on your local machine or in a Docker container (without installing anything beside of Docker).  

### Automated Testing on local machine

* Unit Test `make test`
* Integration Test `make integration`

Test coverage 
```bash
go test --cover ./...
```

### Automated Testing with docker

Full test (Unit + Integration):
1. `docker build --tag jsonrpctest . -f test.dockerfile`
2. `docker run  --rm jsonrpctest:latest`

## Manual Testing

The JSON RPC server and the `Pharmacy.SearchNearestPharmacy` method is exposed on the path `/api`.  
So using port `8080` the endpoint will be `http://localhost:8080/api`.  

### cURL

```cURL
curl --location --request POST 'http://localhost:8080/api' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": 1,
    "jsonrpc": "2.0",
    "method": "Pharmacy.SearchNearestPharmacy",
    "params": [{
        "currentLocation": {
            "latitude": 41.10938993,
            "longitude": 15.0321010
        },
        "range": 10000,
        "limit": 6
    }]
}'
```

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