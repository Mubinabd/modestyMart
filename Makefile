run:
	go mod tidy
	clear
	go run main.go

proto-gen:
	protoc --go_out=./ --go-grpc_out=./ modestyMart_submodule/*.proto

migrate_up:
	migrate -path migrations -database postgres://postgres:postgres@localhost:5432/modestymart?sslmode=disable -verbose up

migrate_down:
	migrate -path migrations -database postgres://postgres:postgres@localhost:5432/modestymart?sslmode=disable -verbose down

migrate_force:
	migrate -path migrations -database postgres://postgres:postgres@localhost:5432/modestymart?sslmode=disable -verbose force 1

migrate_file:
	migrate create -ext sql -dir migrations -seq create_table

swag-gen:
	~/go/bin/swag init -g ./internal/http/api.go -o internal/http/docs force 1

git:
	git add .
	git commit -m "final verse"
	git push 