include .envrc

MIGRATIONS_PATH=./cmd/migrate/migrations

.PHONY: docker-up migrate-create migrate-up migrate-down seed gen-docs

dokcer-up:
	@docker compose --env-file .envrc up --build

migrate-create:
	@migrate create -seq -ext sql -dir ${MIGRATIONS_PATH} ${filter-out $@,${MAKECMDGOALS}}

migrate-up:
	@migrate -path=${MIGRATIONS_PATH} -database=${DB_ADDR} up

migrate-down:
	@migrate -path=${MIGRATIONS_PATH} -database=${DB_ADDR} down ${filter-out $@,${MAKECMDGOALS}}

seed:
	@go run cmd/migrate/seed/main.go

gen-docs:
	@swag init -g ./api/main.go -d cmd,internal && swag fmt