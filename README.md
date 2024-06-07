# Boilerplate API

This is a boilerplate API for a simple CRUD application. It is built using Golang 1.22 and standard net/http package. It uses Swagger for API documentation and Postgres for the database.

## Installation

1. Clone the repository
2. Run `go mod download` to download the dependencies
3. Create a `.env` file in the root directory and add the following environment variables:
```
DB_USER=your_db_user
DB_PASSWORD=your_db_password
DB_NAME=your_db_name
DB_HOST=your_db_host
DB_PORT=your_db_port
```
4. Run `go run main.go` to start the server

## API Documentation

The API documentation is available at `http://localhost:8080/swagger/index.html`
