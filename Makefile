include .env

up:
	@docker compose up

down:
	@docker compose down -v --remove-orphans

background:
	@docker compose up -d --remove-orphans

start_container:
	@docker start ${DB_CONTAINER_NAME}

stop_container:
	@docker stop ${DB_CONTAINER_NAME}

rm_container:
	@docker rm ${DB_CONTAINER_NAME}

postgres:
	@docker run --name ${DB_CONTAINER_NAME} -e POSTGRES_USER=${USER} -e POSTGRES_PASSWORD=${DB_PASSWORD} -p ${DB_PORT}:5432 -d ${DB_IMAGE}

createdb:
	@docker exec -it ${DB_CONTAINER_NAME} createdb --username=${USER} --owner=${USER} ${DB_NAME}

dropdb:
	@docker exec -it ${DB_CONTAINER_NAME} dropdb ${DB_NAME}

migration:
	@migrate create -ext sql -dir db/migration -seq init_schema

migrateup:
	@migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	@migrate -path db/migration -database "$(DB_URL)" -verbose down