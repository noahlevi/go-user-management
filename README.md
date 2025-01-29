# User Management Microservice

This is a simple microservice application written in Go that manages user information. The application provides a RESTful API for creating, retrieving, updating, and deleting user records stored in a PostgreSQL database.

## Table of Contents

- [Features](#features)
- [Technology Stack](#technology-stack)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Usage](#usage)
- [Docker Setup](#docker-setup)
- [API Endpoints](#api-endpoints)
- [Environment Variables](#environment-variables)
- [License](#license)

## Features

- CRUD operations for user management.
- PostgreSQL as the database for persistent storage.
- Docker for containerization.

## Technology Stack

- **Go**: Programming language used for building the service.
- **Gin**: Web framework for building APIs in Go.
- **GORM**: Object Relational Mapping (ORM) for Golang.
- **PostgreSQL**: Relational database management system.
- **Docker**: Containerization platform.

## Project Structure

.
├── db
│   └── connection.go
├── docker-compose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── main.go
├── models
│   └── user.go
├── README.md
└── user
    ├── handler.go
    └── service.go

## Installation

1. **Clone the repository**

```bash
git clone https://github.com/your-username/user-management.git
cd user-management
```

## Install dependencies

2. **Make sure you have Go installed. Then run:**

```bash
go mod tidy
```

## Create a .env file

```bash
HOST=0.0.0.0
PORT=5000
DB_HOST=localhost
DB_USER=gouser
DB_NAME=yourdb
DB_PASSWORD=yourpassword
DB_PORT=5432
```


## Usage

1. **To run the application, you can execute the following command:**
```bash 
go run main.go
```

2. **If using Docker, use the following command:**
```bash 
docker-compose up --build
```

## API Endpoints

`POST /users - Create a new user.`<br>
`GET /users/{id} - Retrieve user details by ID.`<br>
`PUT /users/{id} - Update user details.`<br>
`DELETE /users/{id} - Delete a user.` <br>

## Example Usage

```bash 
curl -X POST http://localhost:5000/users -H "Content-Type: application/json" -d '{"name": "John Doe", "email": "john@example.com"}'
```

```bash 
curl -X GET http://localhost:5000/users/{id}
```

