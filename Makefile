up:
	docker-compose up --build

down:
	docker-compose down

server:
	go run cmd/server/main.go

migration:
	go run cmd/migration/main.go

publish:
	go run cmd/publisher/main.go