# Project Go Gexagonal Boilerplate

This is the backend service for the API project, built with Go following a hexagonal architecture.

## Project Structure

```
api-app/
├── cmd/
│   └── api/
│       └── main.go
├── docs/
├── go.mod
├── go.sum
├── .gitignore
├── README.md
├── internal/
│   ├── adapter/
│   │   ├── handler/
│   │   │   └── user_handler.go
│   │   ├── repository/
│   │   │   └── user_repository.go
│   │   └── router/
│   │       └── router.go
│   ├── bootstrap/
│   │   └── app.go
│   ├── core/
│   │   ├── domain/
│   │   │   └── user.go
│   │   ├── dto/
│   │   │   ├── mapper/
│   │   │   │   └── user_mapper.go
│   │   │   └── user_dto.go
│   │   ├── port/
│   │   │   └── user.go
│   │   └── service/
│   │       └── user_service.go
├── pkg/
│   ├── config/
│   │   └── config.go
│   ├── database/
│   │   └── database.go
│   ├── response/
│   │   └── response.go
│   ├── utils/
│   └── validator/
│       └── validator.go
``` 