postgres:
	docker run -p 5432:5432 --name postgres -e POSTGRES_USE=postgres -e POSTGRES_PASSWORD=123456 -d postgres:15.4
createdb:
	docker exec -it postgres createdb --username=postgres --owner=postgres simple_bank
dropdb:
	docker exec -it postgres dropdb simple_bank
migrateup:
	migrate -path db/migration -database "postgresql://postgres:123456@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://postgres:123456@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migratedown sqlc