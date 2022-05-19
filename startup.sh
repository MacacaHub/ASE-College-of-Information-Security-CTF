#!/bin/bash

export $(grep -v '^#' .env | xargs -d '\n')

docker network inspect $DEFAULT_NETWORK >/dev/null 2>&1 || docker network create $DEFAULT_NETWORK

for docker_compose_path in $(find ./challenges -name docker-compose.yml); do 
    docker-compose -f $docker_compose_path up -d
done

docker-compose -f ctfcli/docker-compose.yml up -d

docker-compose -f ctfd/docker-compose.yml up -d

docker-compose up -d
