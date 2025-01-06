include .env

# Variables
MIGRATION_PATH = ./migrations
DB_DOCKER_YML_PATH = ./local/docker-compose.yml

# =====================Helpers=======================
## help: print this help message
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /'

.PHONY: confirm
confirm: 
	@echo -n 'Are you sure? [y/N] ' && read ans && [ $${ans:-N} = y ]


# =====================Development=======================
## run-api: run the api	
.PHONY: run-api
run-api:
	go run ./cmd/api

## db-up: start the docker container with the postgres database
.PHONY: db-up
db-up:
	docker-compose -f ${DB_DOCKER_YML_PATH} up -d

## db-down: destroy the docker container with the postgres database
.PHONY: db-down
db-down:
	docker-compose -f ${DB_DOCKER_YML_PATH} down

## db-migration-new: create a new pair of migration files (up and down)
.PHONY: db-migration-new
db-migration-new: 
	@echo "Creating new migration..."
	migrate create -seq -ext=.sql -dir=$(MIGRATION_PATH) $(name)

## db-migration-up: migrate the database to the latest version
.PHONY: db-migration-up
db-migration-up: confirm
	@echo "Running migrations..."
	migrate -path=$(MIGRATION_PATH) -database=$(KELODI_DB_DSN) up


# =====================Quality control=======================
## audit: tidy dependencies and format, vet and test all code
.PHONY: audit
audit: 
	@echo "Tidying and verifying module dependencies..."
	go mod tidy
	go mod verify
	@echo "Formatting code..."
	go fmt ./...
	@echo "Vetting code..."
	go vet ./...
	staticcheck ./...
	@echo "Running tests..."
	go test -vet=off ./...


# =====================Build=======================
current_time = $(shell date --iso-8601=seconds)
git_description = $(shell git describe --always --tags --long  --dirty)
linker_flags = '-s -X main.buildTime=${current_time} -X main.version=${git_description}'

## build-api: build the cmd/api application
.PHONY: build-api 
build-api:
	@echo "Building cmd/api..."
	go build -ldflags=${linker_flags} -o=./bin/api ./cmd/api
	GOOS=linux GOARCH=amd64 go build -ldflags=${linker_flags} -o=./bin/linux_amd64/api ./cmd/api


# =====================Production=======================
production_host_ip=${SERVER_IP}

## prod-connect: connect to the production server
.PHONY: prod-connect
prod-connect:
	ssh kelodi@${production_host_ip}

## prod-vps-setup: setup the production server 
.PHONY: prod-vps-setup
prod-vps-setup:
	rsync -rP -e /usr/bin/ssh --delete ./remote/setup root@${production_host_ip}:/root

## prod-deploy-api: deploy the api to production and migrate db to latest
.PHONY: prod-deploy-api
prod-deploy-api:
	rsync -rP -e /usr/bin/ssh --delete ./bin/linux_amd64/api ./migrations kelodi@${production_host_ip}:~/backend/
	ssh -t kelodi@${production_host_ip} 'migrate -path=./backend/migrations -database=$$KELODI_DB_DSN up' 

## prod-configure-api-service: configure the api service on the production server
.PHONY: prod-configure-api-service
prod-configure-api-service:
	rsync -P -e /usr/bin/ssh ./remote/production/api.service kelodi@${production_host_ip}:~
	ssh -t kelodi@${production_host_ip} 'sudo mv ~/api.service /etc/systemd/system/ && sudo systemctl enable api && sudo systemctl restart api'

## prod-configure-caddy: configure caddy on the production server
.PHONY: prod-configure-caddy
prod-configure-caddy:
	rsync -P -e /usr/bin/ssh ./remote/production/Caddyfile kelodi@${production_host_ip}:~
	ssh -t kelodi@${production_host_ip} 'sudo mv ~/Caddyfile /etc/caddy/ && sudo chown caddy:caddy /etc/caddy/Caddyfile && sudo chmod 644 /etc/caddy/Caddyfile && sudo systemctl reload caddy'