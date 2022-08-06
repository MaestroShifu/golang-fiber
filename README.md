# golang-fiber

# Start project
1. Inicializamos un modulo `go mod init <modulename>`.
2. Decargamos las dependencias de nuestro go.mod `go mod tidy` o en su defecto podemos utilizar `go mod download`
3. Para instalar una libreria utilizamos `go get <modulename>`.
4. Sirve para verificar la integridad entre las descargas en disco y las librerias requeridas `go mod verify`.

# Como correr el proyecto
1. Tenemos que instalar todas las librerias de nuestro proyecto que estan en nuestro `go.mod` y `go.sum`.
```bash
make install
```
