postgres:
	docker run --name postgres16 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres16-alpine

createdb: 
	docker exec -t postgres16 createdb --username root --owner root simplebank_db

dropdb:
	docker exec -t postgres16 dropdb simplebank_db

migrateup:
	migrate -path db/migrations -database "${DB_URL}" -verbose up

migratedown:
	migrate -path db/migrations -database "${DB_URL}" -verbose down

new_migration:
	migrate create -ext sql -dir db/migrations -seq ${name}

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	GIN_MODE=release go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/albugowy15/simplebank/db/sqlc Store
	mockgen -package mockwk -destination worker/mock/distributor.go github.com/albugowy15/simplebank/worker TaskDistributor

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
	proto/*.proto


.PHONY: postgres createdb dropdb migrateup migratedown new_migration sqlc test server mock proto
