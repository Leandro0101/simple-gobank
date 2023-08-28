migrate: 
	migrate -path=db/migration -database "postgres://root:toor123@localhost:5432/database?sslmode=disable" -verbose up

migratedown: 
	migrate -path=db/migration -database "postgres://root:toor123@localhost:5432/database?sslmode=disable" -verbose up
up:
	docker-compose up

sqlc: 
	sqlc generate

.PHONY: migrate migratedown up sqlc