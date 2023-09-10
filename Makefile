migrate: 
	migrate -path=db/migration -database "postgres://root:toor123@localhost:5432/database?sslmode=disable" -verbose up

migratedown: 
	migrate -path=db/migration -database "postgres://root:toor123@localhost:5432/database?sslmode=disable" -verbose up
up:
	docker-compose up -d

sqlc: 
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: migrate migratedown up sqlc test