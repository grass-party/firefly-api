#!/bin/sh

docker-compose -f docker-compose-dev.yml down
docker-compose -f docker-compose-dev.yml up -d
docker attach firefly-was
