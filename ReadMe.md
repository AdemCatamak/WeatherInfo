# Weather Info

## Overview
This project is a simple, weather information searching application built using Go. It leverages 'Open Weather' api and is containerized using Docker for easy deployment and testing. 

## Run Application

### DevContainer (Recommended)

If you choose this method, you will need:

- Docker
- An IDE that supports DevContainers (e.g., VS Code, GoLand)

The DevContainer will redirect incoming requests on port 8080 of your local machine to port 80 within the container. This is the port where your application is configured to listen.

### Local Execution

If you choose this method, you'll need:

- Go 1.22 or later version
- An IDE that you love
