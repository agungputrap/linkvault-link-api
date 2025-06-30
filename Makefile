include .env
export $(shell sed 's/=.*//' .env)

APP_NAME=linkvault-api
MIGRATE_PATH=./migrations
DB_URL=postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

run:
	go run cmd/main.go

migrate-up:
	migrate -path $(MIGRATE_PATH) -database "$(DB_URL)" up

migrate-down:
	migrate -path $(MIGRATE_PATH) -database "$(DB_URL)" down

migrate-create:
	@if [ -z "$(name)" ]; then \
		echo "Usage: make migrate-create name=create_users_table"; \
	else \
		migrate create -ext sql -dir $(MIGRATE_PATH) -seq $(name); \
	fi