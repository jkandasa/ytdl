#!/bin/bash

source ./scripts/version.sh

# container registry
REGISTRY='docker.io/jkandasa'
PLATFORMS="linux/arm/v6,linux/arm/v7,linux/arm64,linux/amd64"
IMAGE_TAG=${VERSION}

# build web console
cd web-console
yarn install && CI=false yarn build
cd -

# build and push to docker.io
docker buildx build --push \
  --progress=plain \
  --build-arg=GOPROXY=${GOPROXY} \
  --platform ${PLATFORMS} \
  --file Dockerfile \
  --tag ${REGISTRY}/ytdl:${IMAGE_TAG} .
