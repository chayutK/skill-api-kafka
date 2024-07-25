build:
	docker
	docker-compose build

start:
	docker-compose up

start-build:
	docker-compose up --build

push:
	docker push kracker71/skill-kafka-api:latest
	docker push kracker71/kafka-consumer:latest