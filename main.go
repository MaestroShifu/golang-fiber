package main

import (
	"log"
	"runtime"

	"github.com/MaestroShifu/golang-fiber/src/infrastructure"
)

func main() {
	app, err := infrastructure.StartApp()
	if err != nil {
		log.Fatal(err)
		return
	}
	app.Listen(":3000")
	log.Println("This is the arquitecture of system: ", runtime.GOOS, runtime.GOARCH)
}
