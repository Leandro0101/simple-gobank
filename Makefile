migrateup: 
	migrate -path=db/migration -database "postgres://root:toor123@localhost:5432/database?sslmode=disable" -verbose up

migrateup1: 
	migrate -path=db/migration -database "postgres://root:toor123@localhost:5432/database?sslmode=disable" -verbose up 1

migratedown: 
	migrate -path=db/migration -database "postgres://root:toor123@localhost:5432/database?sslmode=disable" -verbose down

migratedown1: 
	migrate -path=db/migration -database "postgres://root:toor123@localhost:5432/database?sslmode=disable" -verbose down 1

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

.PHONY: migrateup migratedown up sqlc test server mock migratedown1 migrateup1