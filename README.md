# Kelodi Backend

## Description

Kelodi Backend is a RESTful API service for the Kelodi application, built using Go and backed by a PostgreSQL database.

## Features

- User authentication and authorization
- CRUD operations for various resources

## Installation

1. **Clone the repository**:

   ```sh
   git clone https://github.com/DiogoSilvaa/kelodi-backend.git
   cd kelodi-backend
   ```

2. **Install dependencies**:

   ```sh
   go mod tidy
   ```

3. **Set up environment variables**:
   Create a .env file in the root directory and add the necessary environment variables. Refer to `.env.example` for the required variables.

4. **Run application**:

   ```sh
   make run-api
   ```

Additional commands can be found in [Makefile](Makefile)

## Usage

To use the Kelodi Backend, make HTTP requests to the API endpoints.

### Example Endpoints

- `GET /v1/healthcheck`: Get current server status
- `GET /v1/properties`: Get all properties
- `POST /v1/properties`: Create new property

## Dependencies

- Go
- PostgreSQL

## TODO

- Add sqlx
- Add tests
- Add transactions to multi update actions
