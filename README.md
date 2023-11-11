# person-api-echo

## Migrated files from person-api-http-net to echo framework version

### Handler
- login.login()
- person:
    - create()
    - update()
    - delete()
    - getByID()
    - getAll()
- response.responseJSON() (borrar)
- route
    - RoutePerson()
    - RouteLogin()

### Middleware
- Log() (borrar)
- Authentication()
- forbidden()

### CMD
- Register using echo