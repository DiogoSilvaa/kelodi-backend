# Kelodi API

This folder contains the main application code for the Kelodi API. Below is a brief description of the core files and their purposes.

## Files

### [context.go](context.go)

- **Context Management**: Functions to set and get user information in the request context.

### [errors.go](errors.go)

- **Error Handling**: Functions to log errors and send error responses to the client.

### [healthcheck.go](healthcheck.go)

- **Health Check Endpoint**: Handler for the `/v1/healthcheck` endpoint to check the server status.

### [helpers.go](helpers.go)

- **Helper Functions**: Utility functions for reading request parameters, writing JSON responses, and running background tasks.

### [main.go](main.go)

- **Main Application Entry Point**: Initializes the application, loads configuration, connects to the database, and starts the server.

### [middleware.go](middleware.go)

- **Middleware**: Middleware functions for panic recovery, rate limiting, authentication, CORS, and metrics.

### [routes.go](routes.go)

- **Routing**: Defines the routes for the API endpoints and applies middleware.

### [server.go](server.go)

- **Server Management**: Functions to start and gracefully shut down the server.

## Features

### Filtering, Sorting and Pagination

The API supports filtering, sorting and pagination for various endpoints. You can use query parameters to customize the results. Below are the common query parameters:

- **Filtering**: You can filter results based on specific fields. For example, to filter properties by title, description, or location, you can use the following query parameters:

  - `title`: Filter by property title.
  - `description`: Filter by property description.
  - `location`: Filter by property location.

- **Sorting**: You can sort results based on specific fields. Use the `sort` query parameter to specify the field to sort by. You can also specify the sort direction by prefixing the field with a hyphen (`-`) for descending order. For example:

  - `sort=id`: Sort by ID in ascending order.
  - `sort=-id`: Sort by ID in descending order.
  - `sort=title`: Sort by title in ascending order.
  - `sort=-title`: Sort by title in descending order.

- **Pagination**: You can paginate results using the `page` and `page_size` query parameters:
  - `page`: The page number to retrieve (default is 1).
  - `page_size`: The number of items per page (default is 20).

### JSON Handling

- JSON items with null values will be ignored and remain unchanged in the database.

### Rate Limiting

- **Rate Limiting**: The API has built-in rate limiting to prevent abuse. The rate limit can be configured using the limiter settings in the configuration.

### Authentication and Authorization

- **Authentication**: The API uses Bearer tokens for authentication. Endpoints that require authentication will return a `401 Unauthorized` response if the token is missing or invalid.
- **Authorization**: The API uses permissions to control access to certain endpoints. Users must have the necessary permissions to access these endpoints.

### Error Handling

- **Error Handling**: The API provides detailed error responses with appropriate HTTP status codes. Common error responses include `400 Bad Request`, `401 Unauthorized`, `403 Forbidden`, `404 Not Found`, and `500 Internal Server Error`.

### Background Tasks

- **Background Tasks**: The API can run background tasks using the `background` helper function. This is useful for tasks that do not need to block the main request/response cycle.

### Middleware

- **Middleware**: The API uses middleware for various purposes, including panic recovery, rate limiting, authentication, CORS, and metrics.

### Health Check

- **Health Check**: The API provides a health check endpoint at `/v1/healthcheck` to check the server status.

## Usage

To use the Kelodi API, make HTTP requests to the defined endpoints. Below are some example endpoints:

- `GET /v1/healthcheck`: Get current server status
- `GET /v1/properties`: Get all properties
- `POST /v1/properties`: Create a new property

Refer to the main [README.md](../../README.md) for installation and setup instructions.
