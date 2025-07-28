# Projects

The `golang-api` folder contains three golang projects that implements the same API using differents modules.

- `backend-http`: It implements a simple `CRUD` API with golang `http` module.

- `backend-gorilla-mux`: It implements the same `CRUD` API with golang `gorilla mux` module. Additionally we added a simple middleware for authentication.

- `backend-swagger`: It implemnts the same API as in the other projects but here we are using `swagger` to create the API and middlewares. 

## Run the postgres database

Before executing any of the previous project with golang we need to run a `postgres` database using `docker`.

The code has been tested using `postgres 17.1`

- <https://hub.docker.com/layers/library/postgres/17.1/images/sha256-5dee68706b08e5e13db9d9a98981e62eb4a9837187821d9a30aaaf1b9424864f>

The command to run the `postgres` container with `docker` is the one below:

```shell
run docker run --name mypostgres -p 5432:5432 -e POSTGRES_PASSWORD=password -e POSTGRES_USER=postgres -e POSTGRES_DB=mydb -d postgres:17.1
```

**If you do not have docker installed you can use `Rancher Desktop` <https://rancherdesktop.io>**

After running the docker command you can verify if the `postgres` container is actually running:

```shell
docker ps -f name=mypostgres
```

Connect to the `postgres` container:

```shell
docker exec -it mypostgres psql -U postgres -d mydb
```

The following step is to create the `users` table where all the users information will be stored:

```shell
# inside the container run
\c mydb;

# create the table to store the data
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL
);
```

Finally, add some information inside the table:

```shell
INSERT INTO users (name, email) VALUES 
    ('Peter Smith', 'peter.smith@example.com'),
    ('Juan García', 'juan.garcia@example.com'),
    ('Przemek Majewsky', 'przemek.majewsky@example.com');
```

## Executing the projects

### backend-swagger project

We will start with the `backend-swagger` project.

In this project we are defining our API using a `swagger.yaml` file in `SWAGGER 2.0`. You can find the `swagger.yaml` file in `backend-swagger/swagger.yaml`

Once the API is defined in the `swagger.yaml` file we can run the command to generate the following files:

- `cmd/user-management-api-server/main.go` → Access point to the server
- `restapi/` → API configuration and middlewares.
- `restapi/operations/` → empty handlers for the endpoints.
- `models/` → generated models from swagger.yaml.

```shell
swagger generate server --target "./" --name UserManagementAPI --spec "swagger.yaml" --principal "interface{}"
```

**The previous command assumes that you are inside the path `golang-projects/backend-swagger`**

**Note: If you clone the repository all the models and handlers are already created so it is not necessary to run the previous command, this is just informative.**

Before executing the code you must have the `postgres` docker container running with some data already created. 

After that you can execute the code inside `backend-swagger`:

```shell
go run cmd/user-management-api-server/main.go
```

You will see something like this in your terminal

```shell
Serving User Management API at http://localhost:8082
```

* To check if the API is correctly working you can use `Postman` https://www.postman.com and run some curl command

The `test token` that we are using for the authorization is `Secret` (Please, check backend-swagger/utils/authorization.go)

Potsman CURLs:

- Get all users from the database:

```shell
curl -X GET "http://localhost:8082/users" \
  -H "Authorization: Bearer Secret" \
  -H "Content-Type: application/json"
```

**You'll see the index of every user created in the database as well**

- Create a new user in the database:

```shell
curl -X POST "http://localhost:8082/users" \
  -H "Authorization: Bearer Secret" \
  -H "Content-Type: application/json" \
  -d '{
    "id": 1,
    "name": "Juan Pérez",
    "email": "juan@example.com"
}'
```

- Update an existing user:

```shell
curl -X PUT "http://localhost:8082/users/1" \
  -H "Authorization: Bearer Secret" \
  -H "Content-Type: application/json" \
  -d '{
    "id": 1,
    "name": "Juan Actualizado",
    "email": "juan.actualizado@example.com"
}'
```

- Delete a specific user in the database:

```shell
curl -X DELETE "http://localhost:8082/users/1" \
  -H "Authorization: Bearer Secret" \
  -H "Content-Type: application/json"
```

### backend-gorilla-mux project

The steps to run the `postgres` docker container and create the table inside the database with some data are the same as the steps described previously in `backend-swagger project`.

To run the code (supposing you are inside the path golang-projects/backend-gorilla-mux):

```shell
go run main.go
```

You'll see a similar output like the one below on your terminal:

```shell
Connected to the PostgreSQL database
Server is listening on port 8082
```

In postman you can execute the folowung CURLs:

**Note: the token for authorization is `Bearer secret` for the backend-gorilla-mux project**

- Get all users from the database:

```shell
curl -X GET http://localhost:8082/users \
  -H "Authorization: Bearer secret" \
  -H "Content-Type: application/json"
```

- Create a new user in the database:

```shell
curl -X POST "http://localhost:8082/users" \
  -H "Authorization: Bearer secret" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Juan Pérez",
    "email": "juan@example.com"
}'
```

- Update an existing user:

```shell
curl -X PUT "http://localhost:8082/users/1" \
  -H "Authorization: Bearer secret" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Juan Actualizado",
    "email": "juan.actualizado@example.com"
}'
```

- Delete a specific user in the database:

```shell
curl -X DELETE "http://localhost:8082/users/1" \
  -H "Authorization: Bearer secret" \
  -H "Content-Type: application/json"
```

### backend-http project

The steps to run the `postgres` docker container and create the table inside the database with some data are the same as the steps described previously in `backend-swagger project`.

To run the code (supposing you are inside the path golang-projects/backend-http):

```shell
go run main.go
```

```shell
Connected to the PostgreSQL database
Server is listening on port 8082
```

In postman you can execute the folowung CURLs:

**Note: this project has no authorization**

- Get all users from the database:

```shell
curl -X GET http://localhost:8082/users \
  -H "Content-Type: application/json"
```

- Create a new user in the database:

```shell
curl -X POST "http://localhost:8082/users" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Juan Pérez",
    "email": "juan@example.com"
}'
```

- Update an existing user:

```shell
curl -X PUT "http://localhost:8082/users/1" \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Juan Actualizado",
    "email": "juan.actualizado@example.com"
}'
```

- Delete a specific user in the database:

```shell
curl -X DELETE "http://localhost:8082/users/1" \
  -H "Content-Type: application/json"
```