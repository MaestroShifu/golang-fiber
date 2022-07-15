# golang-fiber

# Start project
1. Inicializamos un modulo `go mod init <modulename>`.
2. Decargamos las dependencias de nuestro go.mod `go mod tidy`.
3. Para instalar una libreria utilizamos `go get <modulename>`.
4. Sirve para verificar la integridad entre las descargas en disco y las librerias requeridas `go mod verify`.

# Como correr el proyecto
1. Tenemos que instalar, esto nos servira para poder hacer hot reload en nuestro proyecto. [AIR Github](https://github.com/cosmtrek/air#readme). 
```bash
go install github.com/cosmtrek/air@latest
```
2. Tenemos que instalar todas las librerias de nuestro proyecto que estan en nuestro `go.mod`.
```bash
go mod tidy
```
