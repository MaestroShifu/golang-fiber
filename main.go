package main

import (
	"fmt"
	"runtime"

	"github.com/MaestroShifu/golang-fiber/src/infrastructure"
)

func main() {
	app := infrastructure.StartApp()
	app.Listen(":3000")
	fmt.Println("Esta son las arquitcturas:", runtime.GOOS, runtime.GOARCH)
}
