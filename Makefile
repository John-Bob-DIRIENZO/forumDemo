.PHONY: db db-down migrate migration-down

db:
	docker-compose up -d

db-down:
	docker-compose down

migration:
	docker run -v /Users/jean-francois/go/src/forumDemo/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database postgres://postgres:secret@localhost:5432/postgres?sslmode=disable up 1

migration-down:
	docker run -v /Users/jean-francois/go/src/forumDemo/migrations:/migrations --network host --rm migrate/migrate -path=/migrations/ -database postgres://postgres:secret@localhost:5432/postgres?sslmode=disable down 1
