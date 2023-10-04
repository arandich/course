createdb:
	docker exec -it postgres_container createdb --username=admin bank

dropdb:
	docker exec -it postgres_container dropdb --username=admin bank

migrateup:
	migrate -path db/migration -database "postgresql://admin:admin@localhost:32768/bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://admin:admin@localhost:32768/bank?sslmode=disable" -verbose down

sqlc:
	docker run --rm -v "${CURDIR}:/src" -w /src sqlc/sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb