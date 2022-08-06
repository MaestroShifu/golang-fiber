.PHONY: all # Sirve para poder correr cosas que no son con base a archivos, como comandos de la terminal

IMAGE_NAME := golang_fiber

install:
	go mod download
compile:
	GOOS=freebsd GOARCH=386 go build -o bin/main-freebsd-386 main.go
	GOOS=linux GOARCH=386 go build -o bin/main-linux-386 main.go
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 main.go
build:
	go build -o bin/main main.go
build-image:
	docker build . -t $(IMAGE_NAME)
test:
	go test
dev:
	APP_ENV=development go run main.go
start:
	APP_ENV=production ./bin/main
start-image:
	docker run -d --publish 3000:3000 ${IMAGE_NAME}