BINARY_NAME=lowkey

build:
	go build -o ${BINARY_NAME} cmd/lowkey/main.go

run:
	go run cmd/lowkey/main.go

tidy:
	go mod tidy

