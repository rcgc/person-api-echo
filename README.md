# person-api-echo

## Migrated files from person-api-http-net to echo framework version

### Handler
[x] login.login()
- person:
    [x] create()
    [x] update()
    [x] delete()
    [x] getByID()
    [x] getAll()
[x] response.responseJSON() (borrar)
- route
    [x] RoutePerson()
    [x] RouteLogin()

### Middleware
[x] Log() (borrar)
[x] Authentication()
[x] forbidden() (borrar)

### CMD
[x] Register using echo

### Tests
Tipos de pruebas:<br>
* Tests unitarios: prueban funcionalidades específicas del sistema
* Tests de integración: evalúan la interacción de sistemas externos y que interactúan con el nuestro

Comandos para ejecutar tests en Go:<br>
* All tests in folder: `go test`
* All tests in folder (verbose): `go test -v`
* Execute especific test: `go test TestName`
* Execute especific test (verbose): `go test TestName -v`
* Recursive execution (from root): `go test ./...`

<br>

NOTE: Tests run in concurrent way, therefore, execute `go test ./...` is recommended for unit tests but not for integration tests, in order to solve this last, we could create a script all tests into every package. An example of such script might be written in unix bash format: <br>

#!/bin/sh<br><br>

for d in $(go list ./...); do <br>
    echo "Testeando el paquete $d" <br>
    go test -v $d <br>
done<br>

* Privileges: `ch mod a+x testing.sh`
* Execution: `./testing.sh`