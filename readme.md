# Kasir Sederhana

A CRUD API for managing `Barang` using Go, PostgreSQL, and Golang Migrate.


## Project Structure

```
myapp/
├── main.go
├── config/
│   └── config.go
├── controllers/
│   └── barang.go
├── models/
│   └── barang.go
├── routes/
│   └── routes.go
├── migrations/
│   └── 001_create_barang_table.up.sql
│   └── 001_create_barang_table.down.sql
├── go.mod
└── go.sum
```

## Setup

### Step 1: Initialize Go Module

```sh
go mod init myapp
```

### Step 2: Install Dependencies

```sh
go get -u github.com/jackc/pgx/v4
go get -u github.com/gorilla/mux
go get -u github.com/golang-migrate/migrate/v4
go get -u github.com/joho/godotenv
```

### Step 3: Create `.env` File

Create a `.env` file in the root directory with the following content:

```
DATABASE_URL=postgres://username:password@localhost:5432/mydatabase?sslmode=disable
```

Replace `username`, `password`, and `mydatabase` with your PostgreSQL credentials and database name.

### Step 4: Run Migrations

Install the `migrate` CLI tool and run the migrations:

```sh
go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

migrate -database ${DATABASE_URL} -path ./migrations up
```

### Step 5: Run the Server

```sh
go run main.go
```

## API Endpoints

- **Get all items**: `GET /barang`
- **Get item by ID**: `GET /barang/{id}`
- **Create item**: `POST /barang` with JSON body `{"nama": "ItemName", "harga": 12345}`
- **Update item**: `PUT /barang/{id}` with JSON body `{"nama": "UpdatedName", "harga": 54321}`
- **Delete item**: `DELETE /barang/{id}`
go mod init myapp
