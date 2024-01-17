build:
	go build -o ./bin/main main.go

dev:
	go run main.go

test:
	go test ./

run: test
	@make build
	./bin/main
	
	