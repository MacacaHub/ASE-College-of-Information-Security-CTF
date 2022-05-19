#!/bin/bash

export $(grep -v '^#' .env | xargs -d '\n')

for docker_compose_path in $(find ./challenges -name docker-compose.yml); do 
    docker-compose -f $docker_compose_path down
done
