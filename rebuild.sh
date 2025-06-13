#!/bin/bash
cd /home/ivan/Projects/go/cleanArch
docker-compose down
docker-compose up --build -d
echo "Container reconstruído com sucesso!"
docker logs go-app --tail 10 