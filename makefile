local:
	docker run --net todo-network -p 6094:6093 todo-app

db:
	docker run --name todo-db -e POSTGRES_PASSWORD='qwerty' -p 5436:5432 --net todo-network -d --rm postgres
	
migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disable' up