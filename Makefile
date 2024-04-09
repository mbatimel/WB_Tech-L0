up:
	docker-compose up --build

down:
	docker-compose down

server:
	go run cmd/server/main.go

migrations:
	go run cmd/migration/main.go

publish:
	go run cmd/publish/main.go