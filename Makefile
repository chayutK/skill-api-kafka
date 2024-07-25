build:
	docker compose build

run:
	docker compose up

run-build:
	docker compose up --build

push:
	docker compose build
	docker tag skill-api:latest ghcr.io/chayutk/skill-kafka-api:latest
	docker push ghcr.io/chayutk/skill-api:latest
	docker tag kafka-consumer:latest ghcr.io/chayutk/kafka-consumer:latest
	docker push ghcr.io/chayutk/kafka-consumer:latest 

test:
	cd e2e/ && npx playwright test