# Students API

A RESTful API built with Go for managing student records. This API provides CRUD (Create, Read, Update, Delete) operations for student data with SQLite as the storage backend.

## Features

- Create new students
- Get student by ID
- Get list of all students
- Delete student by ID
- Request validation
- Structured logging
- Graceful server shutdown
- SQLite database storage
- Configuration via YAML file

## Prerequisites

- Go 1.25.1 or higher
- SQLite3 (usually comes pre-installed on most systems)

## Project Structure

```
golang-student/
├── cmd/
│   └── students-api/
│       └── main.go              # Application entry point
├── config/
│   └── local.yaml               # Configuration file
├── internal/
│   ├── config/
│   │   └── config.go            # Configuration management
│   ├── http/
│   │   └── handlers/
│   │       └── student/
│   │           └── student.go   # HTTP handlers for student operations
│   ├── storage/
│   │   ├── sqlite/
│   │   │   └── sqlite.go        # SQLite implementation
│   │   └── storage.go           # Storage interface
│   ├── types/
│   │   └── types.go             # Data types
│   └── utils/
│       └── response/
│           └── response.go      # Response utilities
├── storage/
│   └── storage.db               # SQLite database (created automatically)
├── go.mod
└── README.md
```

## Installation

1. Clone the repository:
```bash
git clone <repository-url>
cd golang-student
```

2. Install dependencies:
```bash
go mod download
```

## Configuration

The application uses a YAML configuration file. An example configuration is provided in `config/local.yaml`:

```yaml
env: "dev"
storage_path: "storage/storage.db"
http_server_port:
  address: "localhost:8082"
```

### Configuration Options

- `env`: Environment name (e.g., "dev", "prod")
- `storage_path`: Path to the SQLite database file
- `http_server_port.address`: Server address and port

You can specify the config file path using:
- Environment variable: `CONFIG_PATH=/path/to/config.yaml`
- Command-line flag: `go run cmd/students-api/main.go -config /path/to/config.yaml`

## Running the Application

1. Make sure you have a configuration file (e.g., `config/local.yaml`)

2. Set the config path (if not using default):
```bash
export CONFIG_PATH=config/local.yaml
```

3. Run the application:
```bash
go run cmd/students-api/main.go
```

Or build and run:
```bash
go build -o bin/students-api cmd/students-api/main.go
./bin/students-api
```

The server will start on the address specified in your configuration file (default: `localhost:8082`).

## API Endpoints

### Create Student

Create a new student record.

**Request:**
```http
POST /api/students
Content-Type: application/json

{
  "name": "John Doe",
  "email": "john.doe@example.com",
  "age": 20
}
```

**Response:**
```json
{
  "Id": 1
}
```

**Status Codes:**
- `201 Created` - Student created successfully
- `400 Bad Request` - Invalid request body or validation error
- `500 Internal Server Error` - Server error

### Get Student by ID

Retrieve a specific student by their ID.

**Request:**
```http
GET /api/students/{id}
```

**Response:**
```json
{
  "id": 1,
  "name": "John Doe",
  "email": "john.doe@example.com",
  "age": 20
}
```

**Status Codes:**
- `200 OK` - Student found
- `400 Bad Request` - Invalid ID format
- `500 Internal Server Error` - Server error

### Get All Students

Retrieve a list of all students.

**Request:**
```http
GET /api/students
```

**Response:**
```json
[
  {
    "id": 1,
    "name": "John Doe",
    "email": "john.doe@example.com",
    "age": 20
  },
  {
    "id": 2,
    "name": "Jane Smith",
    "email": "jane.smith@example.com",
    "age": 22
  }
]
```

**Status Codes:**
- `200 OK` - Success
- `500 Internal Server Error` - Server error

### Delete Student

Delete a specific student by their ID.

**Request:**
```http
DELETE /api/students/{id}
```

**Response:**
```json
{
  "message": "student deleted successfully",
  "id": 1
}
```

**Status Codes:**
- `200 OK` - Student deleted successfully
- `400 Bad Request` - Invalid ID format
- `500 Internal Server Error` - Server error

## Data Model

### Student

```go
type Student struct {
    ID    int64  `json:"id"`
    Name  string `json:"name" validate:"required"`
    Email string `json:"email" validate:"required"`
    Age   int    `json:"age" validate:"required"`
}
```

**Validation Rules:**
- `name`: Required
- `email`: Required
- `age`: Required

## Error Responses

### General Error
```json
{
  "status": "error",
  "error": "error message"
}
```

### Validation Error
```json
{
  "status": "error",
  "error": "field name is required, field email is required"
}
```

## Database Schema

The application automatically creates the following table on startup:

```sql
CREATE TABLE IF NOT EXISTS students (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    age INTEGER NOT NULL
)
```

## Development

### Building

```bash
go build -o bin/students-api cmd/students-api/main.go
```

### Testing

You can test the API using `curl` or any HTTP client:

```bash
# Create a student
curl -X POST http://localhost:8082/api/students \
  -H "Content-Type: application/json" \
  -d '{"name":"John Doe","email":"john@example.com","age":20}'

# Get all students
curl http://localhost:8082/api/students

# Get student by ID
curl http://localhost:8082/api/students/1

# Delete student
curl -X DELETE http://localhost:8082/api/students/1
```

## Dependencies

- [go-playground/validator](https://github.com/go-playground/validator) - Request validation
- [mattn/go-sqlite3](https://github.com/mattn/go-sqlite3) - SQLite driver
- [ilyakaznacheev/cleanenv](https://github.com/ilyakaznacheev/cleanenv) - Configuration management

## Author

Manish kumar

