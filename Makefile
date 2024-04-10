.PHONY: server
server:
	go build -v ./cmd/server
	./server

.PHONY: migration
migration:
	go build -v ./cmd/migration
	./migration

.PHONY: test
test:
	echo "GET http://localhost:1234/order/" | vegeta attack -duration=10s | tee results.bin | vegeta report

.PHONY: publisher
publisher:
	go build -v ./cmd/publisher

.PHONY: up
up:
	docker-compose up -d

.PHONY: down
down:
	docker-compose down

.DEFAULT_GOAL := build