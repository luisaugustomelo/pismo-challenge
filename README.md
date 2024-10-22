# pismo-challenge

## Project Setup Instructions

This project consists of a backend (written in Go). The backend relies on a database that can be easily started using Docker. Follow the instructions below to get the entire project up and running.

Prerequisites

Docker and Docker Compose installed
Go installed (version 1.18 or higher)

### Step-by-Step Setup

#### 1. Start the Database with Docker
To start the application environment (install dependencies, docker-compose up and swagger doc generation):

```bash
make run
```

Starting application:

```
go run main.go
```

## Environment Variables Setup

This project requires environment variables to be set for proper database connections and other configurations. You can provide these in a `.env` file in the project root.

### 1. Create a `.env` file

Copy the provided `.env-example` file and rename it to `.env`. Update the values in the `.env` file according to your environment setup.

Example:
```bash
cp .env-example .env
```
Update the .env file with your environment-specific details.

#### Fallback to .env-example
If the .env file is not present, the application will automatically use the values from .env-example.

### Swagger Documentation
To generate and serve Swagger documentation for the API, follow these steps:

Generate Swagger docs:
```
make swag
```

Access the Swagger documentation at http://localhost:3333/swagger/index.html after starting the server.

## Accessing the Application

Backend: The backend API will be running on http://localhost:3333.

### Summary of Commands
```
# Start all environment
make docker-up

# Run the Go backend
go run main.go

# Drop database tables
make migrate-drop

# Run tests
make test

# Generate Swagger documentation
make swag

# Start the environment
make run

# Stop the environment
make stop

# Stop and remove Docker containers
make docker-down
```

## Troubleshooting

- If there are issues with Docker, ensure that Docker is properly installed and running.
- Make sure the backend and frontend are started in the correct directories.
- If ports 3333 is already in use, you may need to stop other services or change the ports in the respective configurations.

## Additional Information

This project includes a Makefile for simplifying common tasks such as starting Docker containers, running database migrations, generating Swagger documentation, and running tests. Below is a brief description of the available Makefile commands:

- make run: Starts the database, generates Swagger docs, installs Go dependencies, and runs migrations.
- make stop: Stops all services and removes networks created by Docker Compose.
- make swag: Generates the Swagger documentation.
- make deps: Installs the Go project dependencies.
- make migrate-drop: Drops all database tables.
- make test: Runs all the tests in the project.
- make docker-down: Stops and removes Docker Compose containers and networks.