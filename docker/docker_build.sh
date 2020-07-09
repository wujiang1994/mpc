#!/usr/bin/env bash

set -e
set -o pipefail
PROJECT_NAME="mpc"
MODULE="test"
IMAGE_NAME=${PROJECT_NAME}_${MODULE}
REGISTRY_HOST="127.0.0.1:5000"

# Get the version of current image
echo [Notice]Please input the version of ${IMAGE_NAME}, for example "1.0.0":
read VERSION

echo Begin to update dependences by go module
export GO111MODULE=on
go mod tidy
go mod vendor

echo Begin to build docker image ${IMAGE_NAME}:${VERSION}
echo The enviroment parameter GOPATH is ${GOPATH}

# context
docker build -t ${REGISTRY_HOST}/${IMAGE_NAME}:${VERSION} -f ./Dockerfile ../

# delete build
docker image prune -f

# run local
docker run -d -p 5000:5000 --name=${IMAGE_NAME} ${REGISTRY_HOST}/${IMAGE_NAME}:${VERSION}

# echo logs
# docker logs -f ${IMAGE_NAME}

# docker push ${REGISTRY_HOST}/${IMAGE_NAME}:${VERSION}

# kubectl replace -f ../deployments/manager_deployment.yaml

# kubectl get pods -n member-center | grep manager