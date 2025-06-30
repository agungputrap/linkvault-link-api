# 🔗 LinkVault – Link Management API

**LinkVault** is a simple, lovable, and complete personal project designed as a link management system. It allows users to register, store, organize, and manage their personal links with tags.

This backend service is built with **Go** using the **Fiber framework**, applies **Domain-Driven Design (DDD)**, and connects to a **PostgreSQL** database.

---

## Tech Stack

- **Go Fiber** – web framework
- **GORM** – ORM for PostgreSQL
- **JWT Auth** – for secure user authentication
- **PostgreSQL** – relational database
- **golang-migrate** – database migration tool
- **Domain-Driven Design (DDD)** – modular project structure

---

## Features

- User registration & login with JWT
- Create, update, delete, and list personal links
- Tag support: list tags by user
- Secure route protection with JWT middleware
- Clean architecture with DDD

---

## Project Structure (Simplified)
```bash
linkvault-link-api/
├── cmd/                  # main.go entrypoint
├── internal/
│   ├── domain/           # core entities (user, link, tag)
│   ├── application/      # usecases and DTOs
│   ├── interfaces/       # HTTP handlers and routes
│   ├── infrastructure/   # database, external implementations
│   └── migrations/       # SQL files for database schema
└── go.mod / go.sum       # Go dependencies
```
---

## Getting Started

### 1. Clone & Install

```bash
git clone https://github.com/agungputrap/linkvault-link-api.git
cd linkvault-link-api
go mod tidy
```

### 2. Setup Environment
Create a .env file:
```bash
PORT=3000
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=linkvault
JWT_SECRET=your_jwt_secret
```

### 3. Run Database Migration
Make sure PostgreSQL is running, then:
```bash
make migrate-up
```

### 4. Run app
```bash
make run
```
---
## Personal Note

This project is part of my personal portfolio to:
- Sharpen my skills in Go, backend development, and architecture 
- Practice building microservices and APIs 
- Prepare for freelance and startup opportunities