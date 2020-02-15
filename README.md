# Test Closest Pharmacies

## How to start the application

### Requirements
* Docker

### Steps

1. `docker run`


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