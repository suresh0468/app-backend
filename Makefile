postgres:
	docker run --name gita-postgres -p 5432:5432 -e POSTGRES_USER=suresh -e POSTGRES_PASSWORD=2003 -d postgres:18-alpine

postgresStart:
	docker start gita-postgres

postgresStop:
	docker stop gita-postgres

createdb: 
	docker exec -it gita-postgres createdb --username=suresh --owner=suresh gita_db

migrateUp:
	migrate -path db/migrations -database "postgresql://suresh:2003@localhost:5432/gita_db?sslmode=disable" -verbose up

migrateDown:
	migrate -path db/migrations -database "postgresql://suresh:2003@localhost:5432/gita_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

server:
	go run main.go

.PHONY: postgres postgresStart postgresStop createdb migrateUp migrateDown sqlc server
