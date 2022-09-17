compose_file := "./docker/docker-compose.yaml"
project_name := "bookshelf"
migrate = ""
MARIADB_URL="mysql://golang:golang@tcp(localhost:33061)/bookshelf?parseTime=true"

up:
	docker-compose -f $(compose_file) -p $(project_name) up -d

down:
	docker-compose -f $(compose_file) -p $(project_name) down

down/v:
	docker-compose -f $(compose_file) -p $(project_name) down -v

migrate/create:
	migrate create -ext sql -dir database/migrations -seq $(migrate)

migrate/up:
	migrate -database $(MARIADB_URL) -path database/migrations up

migrate/down:
	migrate -database $(MARIADB_URL) -path database/migrations down
