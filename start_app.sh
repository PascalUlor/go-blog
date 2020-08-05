#!/bin/bash

docker-compose -f ./.dockerApp/docker-compose.yml up --build --remove-orphans -d

