build:
	go build -o ./bin/friends main.go

dev:
	go run main.go

test:
	go test ./

run: test
	@make build
	./bin/friends
	
	