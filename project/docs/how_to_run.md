
# How To Run


## Prerequisites

- Install [Go](https://golang.org/doc/install)


- Install [gomock](https://github.com/golang/mock)


- Install [go-migrate](https://github.com/DavidHuie/gomigrate/blob/master/README.md)


- Install `protoc-gen-go-grpc` & `protoc-gen-go` 

    ```
    $ go install \
        google.golang.org/protobuf/cmd/protoc-gen-go \
        google.golang.org/grpc/cmd/protoc-gen-go-grpc
    ```

    then make sure that these binaries are generated in $GOBIN;

    - `protoc-gen-go-grpc`
    - `protoc-gen-go`

    Make sure that $GOBIN is in $PATH.

- PostgreSQL

    Follow [PostgreSQL download](https://www.postgresql.org/download/).
    
## Run Instruction    

### Using Docker
 ```
    make docker-compose
```

### Manual
- copy env config from env.example
 ```
    cp env.example .env
```

- set env variables with prefix POSTGRES_ according to your Postgres settings

- download depedencies
```
    make tidy
```

- run 
```
    make run
```