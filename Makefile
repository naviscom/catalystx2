postgres:
	docker run --name postgres --network naviscom-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret999 -d postgres:latest

createdb:
	docker exec -it postgres createdb --username=root --owner=root catalystx2

dropdb:
	docker exec -it postgres dropdb catalystx2

migrateup:
	migrate -path db/migration -database "postgresql://root:secret999@localhost:5432/catalystx2?sslmode=disable" -verbose up

migrateupaws:
	migrate -path db/migration -database "postgresql://root:secret999@catalyst.cadqx57sergj.us-west-2.rds.amazonaws.com:5432/catalystx2" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret999@localhost:5432/catalystx2?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

runcon:
	docker run --name catalystx2 --network naviscom-network -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret999@postgres:5432/catalystx2?sslmode=disable" catalystx2:latest

.PHONY: postgres createdb dropdb migrateup migrateupaws migratedown sqlc test server runcon

