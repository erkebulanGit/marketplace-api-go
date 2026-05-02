DB_URL=postgres://postgres:postgres@localhost:5432/marketplace?sslmode=disable

migrate-up:
	migrate -path migrations -database "$(DB_URL)" up

migrate-down:
	migrate -path migrations -database "$(DB_URL)" down

migrate-down-one:
	migrate -path migrations -database "$(DB_URL)" down 1

migrate-version:
	migrate -path migrations -database "$(DB_URL)" version

migrate-create:
	migrate create -ext sql -dir migrations -seq $(NAME)

run:
	go run main.go

deps:
	go mod tidy


docker-build:
	docker-compose build

docker-up:
	docker-compose up

docker-up-d:
	docker-compose up -d

docker-down:
	docker-compose down

docker-logs:
	docker-compose logs -f

docker-ps:
	docker-compose ps

docker-migrate-up:
	docker-compose exec marketplace-api migrate -path /app/migrations -database "postgres://postgres:postgres@postgres:5432/marketplace?sslmode=disable" up