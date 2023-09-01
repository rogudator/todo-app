include .env
export

local:
	go run cmd/main.go

local_db:
	docker run --name ${DB_NAME} -e POSTGRES_PASSWORD='${DB_PASSWORD}' -p ${DB_PORT}:5432 -d --rm postgres
	
local_migrate:
	migrate -path ./schema -database 'postgres://${DB_NAME}:${DB_PASSWORD}@localhost:${DB_PORT}/postgres?sslmode=${DB_SSLMODE}' up