#!/usr/bin/env bash

docker-compose run backend-service glide install
docker-compose run backend-service go build main.go