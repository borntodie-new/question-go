# build a new postgres container
postgres_up:
	@echo "build a new postgres container..."
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:14.0
	@echo "buold a new postgres container success!"
# delete exist container of postgres
postgres_down:
	@echo "stop postgres container..."
	docker stop postgres
	@echo "delete postgres container..."
	docker rm postgres
	@echo "delete postgres container success..."
# create database
create_db:
	@echo "create database..."
	docker exec -it postgres createdb --username=root --owner=root question
	@echo "create database done..."
# drop database
drop_db:
	@echo "drop database..."
	docker exec -it postgres dropdb question
	@echo "drop database done..."

# migrate database
migration_up:
	@echo "migrate database for create..."
	migrate -path ./migrations -database "postgres://root:password@localhost:5432/question?sslmode=disable" -verbose up
	@echo "migrate database success..."

migration_down:
	@echo "migrate database for delete..."
	migrate -path ./migrations -database "postgres://root:password@localhost:5432/question?sslmode=disable" -verbose down
	@echo "migrate database success..."
sqlc:
	sqlc generate