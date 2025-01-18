#!/bin/bash

# Set the Docker Hub username and repository name
DOCKER_USERNAME="atha02"
DOCKER_REPO="web-server-jarkom"
IMAGE_TAG="1.0.0"  # You can change this to a version tag if needed

# Build the Docker image
echo "Building the Docker image..."
docker build -t $DOCKER_USERNAME/$DOCKER_REPO:$IMAGE_TAG .

# Log in to Docker Hub (this will prompt for your Docker Hub credentials)
echo "Logging in to Docker Hub..."
docker login

# Push the Docker image to Docker Hub
echo "Pushing the image to Docker Hub..."
docker push $DOCKER_USERNAME/$DOCKER_REPO:$IMAGE_TAG

echo "Build and push completed successfully."
