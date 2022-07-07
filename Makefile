install:
	go mod tidy
compile:
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 src/main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 src/main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 src/main.go
build:
	go build -o bin/main src/main.go
test:
	go test
run:
	go run src/main.go
all:
	echo "Ejecuta un flujo de tareas build run"