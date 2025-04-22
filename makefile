include $(PWD)/.env

docker-run:
	docker rm -f $(PROJECT_NAME)
	docker build -t $(PROJECT_NAME)-image -f deployment/docker/Dockerfile .
	docker run --name $(PROJECT_NAME) --env-file=".env" -d -p $(PORT):$(PORT) $(PROJECT_NAME)-image

docker-compose-run:
	docker compose -f deployment/docker/docker-compose.yml --env-file=.env rm --force --stop
	docker compose -f deployment/docker/docker-compose.yml --env-file=.env --env-file=.env.credentials up --build --detach
