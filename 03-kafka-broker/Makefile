# Define the docker-compose file
COMPOSE_FILE="./docker-compose-haproxy.yml"
COMPOSE_FILE_DC1="./kafka_cluster1/docker-compose.yml"
COMPOSE_FILE_DC2="./kafka_cluster2/docker-compose.yml"

# Define the Docker network name
NETWORK_NAME=kafka_cluster

# Target to create the Docker network if it doesn't already exist
create-network:
	docker network inspect $(NETWORK_NAME) >/dev/null 2>&1 || docker network create -d bridge --subnet=192.168.1.0/28 $(NETWORK_NAME)
	docker network ls
	ip addr show

# Target to start the services defined in the docker-compose file
up: create-network
	docker-compose -f $(COMPOSE_FILE) up -d
	docker-compose -f $(COMPOSE_FILE_DC1) up -d
	docker-compose -f $(COMPOSE_FILE_DC2) up -d

# Target to stop the services
down:
	docker-compose -f $(COMPOSE_FILE) down -v
	docker-compose -f $(COMPOSE_FILE_DC1) down -v
	docker-compose -f $(COMPOSE_FILE_DC2) down -v

ps:
	docker-compose -f $(COMPOSE_FILE) ps
	docker-compose -f $(COMPOSE_FILE_DC1) ps
	docker-compose -f $(COMPOSE_FILE_DC2) ps

# Target to restart the services
restart: down up

# Target to view the logs
logs:
	docker-compose -f $(COMPOSE_FILE) logs -f

# Target to build or rebuild services
build:
	docker-compose -f $(COMPOSE_FILE) build

# Target to stop and remove all containers, networks, volumes, and images created by up
clean: down
	docker network rm $(NETWORK_NAME)

# Help target to display available commands
help:
	@echo "Makefile for managing Docker Compose"
	@echo ""
	@echo "Usage:"
	@echo "  make up           Start the services"
	@echo "  make down         Stop the services"
	@echo "  make restart      Restart the services"
	@echo "  make logs         View logs"
	@echo "  make build        Build or rebuild services"
	@echo "  make clean        Stop and remove all containers, networks, volumes, and images"
	@echo "  make help         Display this help message"

