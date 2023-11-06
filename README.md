# Summary

This repo is a template for projects setting up a backend service in golang that implements a grpc service. It requires a database connection be established. Currently it is setup to access a postgres database on localhost:5432.

# Set-up Steps

1. Make sure docker is installed and the docker engine is running

2. Run the following to pull the database (postgresql) image 
```shell
docker pull postgres
```

3. Run the following to start a docker container on your local machine with postgres running, and portforward for local access.
    - change the <project-name> as needed
    - change the database password as needed (DB_PASSWORD will need to be change in the projects .env as well)
    - change the port as needed (DB_PORT will need to be change in the projects .env as well)
```shell
docker run --name <project-name>-db -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres
```

4. Start the project in the devcontainer included

If followed correctly, the postgres container will be your 'local' db for the projects devcontainer.

5. Change the query and schema files in './pkg/db/sqlc-gen' directory.

6. Run the following in the devcontainer terminal, at the root directory, to generate the models and queries in go
```shell
sh ./bin/sqlc-gen
```

7. Change the grpc methods, services and messages accordingly.

8. Run the following in the devcontainer terminal, at the root directory, to generate protobuf definitions
```shell
sh ./bin/proto-gen
```

9. Test your database connection/setup with the app's built in CLI
```shell
go run main.go migrate ping
```

10. Change the initial migration files in './pkg/db/migration/migrations' directory.

11. If you're ready to populate tables into your postgres instance you can use the app's built in CLI from the devcontainer terminal.

```shell
go run main.go migrate up
```

# Scripts
## Test DB Connection (Ping)
```shell
go run main.go migrate ping
```

## Migrate Up
```shell
go run main.go migrate up
```

## Migrate Down
```shell
go run main.go migrate down
```

## Start GRPC Server
```shell
go run main.go grpc
```

# GRPC Server

After starting the grpc server you can use grpcurl to test (locally or in the devcontainer)
```shell
grpcurl -plaintext localhost:51001 Unit/GetEventList
```

Alternatively you can use grpcui to test (locally or in the devcontainer)
```shell
grpcui -plaintext localhost:51001
```