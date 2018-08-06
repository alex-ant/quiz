#!/bin/bash

export APP_NAME=fe

# Build frontend.
npm run build

# Build docker image.
TAG=quiz_$APP_NAME:latest
docker build -f Dockerfile -t $TAG .
