compose_file := "./docker/docker-compose.yaml"
project_name := "bookshelf"
migrate = ""
POSTGRESQL_URL := "postgres://postgres:postgres@localhost:54321/postgres?sslmode=disable"

up:
	docker-compose -f $(compose_file) -p $(project_name) up -d

down:
	docker-compose -f $(compose_file) -p $(project_name) down

down/v:
	docker-compose -f $(compose_file) -p $(project_name) down -v

migrate/create:
	migrate create -ext sql -dir database/migrations -seq $(migrate)

migrate/up:
	migrate -database $(POSTGRESQL_URL) -path database/migrations up

migrate/down:
	migrate -database $(POSTGRESQL_URL) -path database/migrations down
