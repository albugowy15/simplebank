postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres16-alpine

createdb: 
	docker exec -t postgres16 createdb --username root --owner root simplebank_db

dropdb:
	docker exec -t postgres16 dropdb simplebank_db

migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simplebank_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simplebank_db?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown test server
