build:
	docker-compose build
.PHONY: build

compose-up:
	docker-compose up
.PHONY: compose-up

compose-down:
	docker-compose down --remove-orphans
.PHONY: compose-down

migrate-up:
	docker-compose run --rm migrate -path=/migrations -database="postgres://postgres:admin@db:5432/db_1?sslmode=disable" up
.PHONY: migrate-up

migrate-down:
	docker-compose run --rm migrate -path=/migrations -database="postgres://postgres:admin@db:5432/db_1?sslmode=disable" down
.PHONY: migrate-down
