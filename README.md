# URL Shortener

A URL shortening service built with Go and PostgreSQL. This project focuses on backend development practices, database integration, and maintainable code organization.

## Tech Stack

- **Language:** Go 1.23.3
- **Database:** PostgreSQL 17
- **Database Queries:** SQLC
- **Containerization:** Docker Compose
- **Configuration:** Environment variables with `.env`
- **Dependencies:** `lib/pq`, `godotenv`, `google/uuid`

## Features

- Short URL generation and management
- PostgreSQL integration for persistent data storage
- SQL query generation with SQLC
- Environment-based configuration
- Docker Compose setup for local database development

*Note: Confirm that each listed application feature is implemented before treating this section as a final feature list.*

## Project Structure

```text
url-shortener/
├── cmd/
│   └── server/       # Application entry point
├── db/
│   ├── migrations/   # Database migrations
│   └── queries/      # SQL queries for SQLC
├── internal/         # Internal application packages
├── .env.example      # Environment variable template
├── compose.yml       # PostgreSQL service configuration
├── sqlc.yaml         # SQLC configuration
├── go.mod
└── go.sum
```

## Getting Started

### Prerequisites

Make sure you have the following installed:

- Go 1.23.3 or later
- Docker and Docker Compose
- Git

### 1. Clone the Repository

```bash
git clone https://github.com/Permanaa/url-shortener.git
cd url-shortener
```

### 2. Configure Environment Variables

Create a `.env` file from the example:

```bash
cp .env.example .env
```

Configure your database credentials in `.env`:

```env
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=url_shortener
DB_HOST=localhost
DB_PORT=5432
```

Replace `your_password` with a secure local development password.

### 3. Start PostgreSQL

Start the database using Docker Compose:

```bash
docker compose up -d
```

Check the database container status:

```bash
docker compose ps
```

### 4. Install Dependencies

```bash
go mod download
```

### 5. Run Database Migrations

Apply the SQL migration files in `db/migrations` using the migration tool configured for this project.

### 6. Run the Application

```bash
go run ./cmd/server
```

The application should start using the configuration defined in your environment variables.

## Database Query Generation

This project uses SQLC to generate type-safe Go code from SQL queries.

After modifying SQL queries or database schemas, regenerate the database code using:

```bash
sqlc generate
```

The generated code is configured to be placed in `internal/dbgen`.

## API Documentation

Base path: `/shortnd`

### Endpoints

| Method | Endpoint | Description |
|---|---|---|
| `POST` | `/shortnd` | Create a shortened URL |
| `GET` | `/shortnd/{shortCode}` | Retrieve a shortened URL by code or perform its configured redirect behavior |
| `PUT` | `/shortnd/{shortCode}` | Update an existing shortened URL |
| `DELETE` | `/shortnd/{shortCode}` | Delete a shortened URL |
| `GET` | `/shortnd/{shortCode}/stats` | Retrieve statistics for a shortened URL |

### 1. Create a Short URL

Create a new shortened URL.

**Request**

```http
POST /shortnd
Content-Type: application/json
```

```json
{
  "url": "https://example.com",
  "shortChode": "customLink"
}
```

**Expected behavior**
- Accepts the original URL.
- Optional shortCode for custom link.
- Generates or assigns a short code.
- Stores the URL mapping in the database.

### 2. Get a Short URL

Retrieve a shortened URL using its short code.

**Request**

```http
GET /shortnd/{shortCode}
```

Example:

```http
GET /shortnd/Ab12Cd
```

**Expected behavior**
- Looks up the short code.
- Returns the corresponding URL or performs a redirect, depending on the handler implementation.

### 3. Update a Short URL

Update the destination of an existing short URL.

**Request**

```http
PUT /shortnd/{shortCode}
Content-Type: application/json
```

```json
{
  "url": "https://example.com/updated"
}
```

**Expected behavior**
- Finds the existing short code.
- Updates its destination URL.
- Persists the changes in the database.

### 4. Delete a Short URL

Delete an existing shortened URL.

**Request**

```http
DELETE /shortnd/{shortCode}
```

Example:

```http
DELETE /shortnd/Ab12Cd
```

**Expected behavior**
- Finds the URL mapping.
- Removes the mapping from the database.

### 5. Get URL Statistics

Retrieve statistics associated with a shortened URL.

**Request**

```http
GET /shortnd/{shortCode}/stats
```

Example:

```http
GET /shortnd/Ab12Cd/stats
```

**Expected behavior**
- Looks up the short code.
- Returns the statistics tracked by the application, such as click counts, if implemented.
