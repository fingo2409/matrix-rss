#!/bin/bash

#MATRIX_RSS_VERSION=$(cat MATRIX_RSS_VERSION)
MATRIX_RSS_VERSION="latest"
IMAGE="ghcr.io/fingo2409/matrix-rss:${MATRIX_RSS_VERSION}"

docker build \
    --build-arg MATRIX_RSS_VERSION="${MATRIX_RSS_VERSION}" \
    --build-arg IMAGE="${IMAGE}" \
    --build-arg JOBS=$(nproc) \
    --no-cache \
    --progress=auto \
    -t "${IMAGE}" .
