# backend-api-go
A simple RESTful API (CRUD) built with GoLang

## Requirements

- [Git](https://git-scm.com/)
- [Docker](https://www.docker.com/)
- [docker-compose](https://docs.docker.com/compose/)

## Install
Clone and setup:

```bash
$ git clone https://github.com/wilk/backend-api-go.git
$ cd backend-api-go
$ ./setup.sh
```

## Running
First, setup docker-compose:

```bash
$ docker-compose up
```

Then start querying at `http://localhost:3000/api/users/`

## Shutdown
```bash
$ docker-compose down
```

## APIs
The entity **User** has the following fields:

- ID (uint)
- name (string)
- email (string)
- age (int)
- mobile (string)

Follows the list of users APIs:

|METHOD|URL|REQUEST HEADERS|REQUEST PAYLOAD|RESPONSE HEADERS|RESPONSE PAYLOAD|
|------|---|---------------|---------------|----------------|----------------|
|GET|http://localhost:3000/api/users/ | | | |User[]|
|POST|http://localhost:3000/api/users/ |Content-Type: "application/json"|User||User|
|GET|http://localhost:3000/api/users/10 | | | |User|
|PUT|http://localhost:3000/api/users/10 |Content-Type: "application/json"|User||User|
|DELETE|http://localhost:3000/api/users/10 | | | | |
