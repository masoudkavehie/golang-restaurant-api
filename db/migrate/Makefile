postgres:
	docker run --name postgres-returant -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine
createdb:
	docker exec -it postgres-returant createdb --username=root --owner=root resturant

dropdb:
	docker exec -it postgres-returant dropdb resturant


migrateup:

	migrate -path . -database "postgresql://root:secret@localhost:5432/resturant?sslmode=disable" --verbose up


migratedown:

	migrate -path . -database "postgresql://root:secret@localhost:5432/resturant?sslmode=disable" --verbose down


sqlc:
	sqlc generate


.PHONY: createdb postgres dropdb migratedown migrateup sqlc
