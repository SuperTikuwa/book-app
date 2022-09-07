compose_file := "./docker/docker-compose.yaml"
project_name := "bookshelf"

up:
	docker-compose -f $(compose_file) -p $(project_name) up -d

down:
	docker-compose -f $(compose_file) -p $(project_name) down

down/v:
	docker-compose -f $(compose_file) -p $(project_name) down -v