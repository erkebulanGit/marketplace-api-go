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