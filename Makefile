# Define variables
DOCKER_COMPOSE_FILE = docker-compose.yml
MIGRATION_FILE = ./migrations/migration_routine.go

# Default targets
.PHONY: all
all: build run

# Starts the development environment
.PHONY: run
run: docker-up swag deps migrate ## Starts docker-compose and runs migrations

# Stops the development environment
.PHONY: stop
stop: docker-down ## Stops containers and removes networks

# Docker compose
.PHONY: docker-up
docker-up: ## Starts the environment with docker-compose
	@echo "Starting docker-compose..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

.PHONY: docker-down
docker-down: ## Stops and removes containers, networks, and volumes created by docker-compose
	@echo "Stopping and removing docker-compose services..."
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

# Swagger documentation
.PHONY: swag
swag: ## Generates Swagger documentation
	@echo "Generating Swagger docs..."
	swag init

# Go dependencies installation
.PHONY: deps
deps: ## Installs Go dependencies
	@echo "Running go mod tidy to install dependencies..."
	go mod tidy

# Migrations to set up the database
.PHONY: migrate
migrate: ## Runs migrations to create the database tables
	@echo "Running database migrations..."
	go run $(MIGRATION_FILE)

# Migrations to drop the database
.PHONY: migrate-drop
migrate-drop: ## Runs migrations to drop the database tables
	@echo "Dropping tables from database..."
	go run $(MIGRATION_FILE) drop

# Tests
.PHONY: test
test: ## Runs all project tests
	@echo "Running tests..."
	go test ./...

# Clean up
.PHONY: clean
clean: docker-down ## Stops containers, removes networks, and performs general clean-up
	@echo "Cleaning generated files..."
	rm -rf swagger docs

# Help
.PHONY: help
help: ## Displays help for available commands
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-15s\033[0m %s\n", $$1, $$2}'
