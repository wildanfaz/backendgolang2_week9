# Backend Golang 2 Week 9

Vehicle Rental server-side part 2

## Clone
```bash
git clone https://github.com/wildanfaz/backendgolang2_week9.git
```

## Dependencies
```bash
go mod tidy
```

## Features

- CRUD
- Search Query
- Hashing Password (bcrypt)
- JWT
- Authentication
- Authorization
- Cobra Command
- DB Migration

## DB Migration
Migrate
-
```bash
go run main.go migrate --up
```

Rollback
-
```bash
go run main.go migrate --down
```

## Tech Stack

**Client:** -

**Server:** Golang, PostgreSQL, GORM, Gorilla/mux

## Author

- [@wildanfaz](https://www.github.com/wildanfaz)
