.PHONY: all # Sirve para poder correr cosas que no son con base a archivos, como comandos de la terminal

install:
	go mod tidy
compile:
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 src/main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 src/main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 src/main.go
build:
	go build -o bin/main src/main.go
build_air:
	go build -o tmp/main src/main.go
test:
	go test
dev:
	APP_ENV=development air
start:
	APP_ENV=production ./bin/main
all:
	echo "Ejecuta un flujo de tareas build run"