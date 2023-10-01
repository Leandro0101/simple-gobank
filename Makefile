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

server:
	go run main.go

mock:
	mockgen --package mockdb --destination db/mock/store.go simple-gobank/db/sqlc Store

.PHONY: migrate migratedown up sqlc test server mock