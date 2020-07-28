#!/usr/bin/env bash

# Delete every Docker containers
# Must be run first because images are attached to containers
docker rm -f $(docker ps -a -q)

# Delete every Docker image
docker rmi -f $(docker images -q)

versionTag=$(git describe --tags `git rev-list --tags --max-count=1`)
echo $versionTag
dockerVersion=$(echo $versionTag | cut -c2-)
echo $dockerVersion
# HEAD build: commit=$(git rev-parse --short HEAD 2>/dev/null)
docker build --rm --build-arg "commit=tags/${versionTag}" --build-arg "version=${versionTag}" -t "skillzint/nakama:${dockerVersion}" .

imageID=$(docker images --filter=reference="skillzint/nakama:${dockerVersion}" --format "{{.ID}}")
echo $imageID

docker tag "${imageID}" "skillzint/nakama:latest"
docker push "skillzint/nakama:${dockerVersion}"
docker push "skillzint/nakama:latest"
