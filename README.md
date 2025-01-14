# Author
Name: Andreas Nilsson <br>
Email: andreas.nilsson@lnu.se

# REST API with GraphQL Support

This project is a RESTful API implemented with Go, including GraphQL capabilities. It uses Postgres for data storage and is containerized using Docker. 
It provides authentication using JWT.
<br>
### 👉 Note: Only GET, POST is implemented. PUT and DELETE for updating and deleting a device Tbi.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
  - [Clone the Repository](#clone-the-repository)
  - [Environment Setup](#environment-setup)
  - [Build and Run with Docker](#build-and-run-with-docker)
- [API Endpoints](#api-endpoints)
  - [Authentication](#authentication)
  - [Devices](#devices)
- [GraphQL Query Example](#graphql-query-example)
- [Design Patterns and Principles](#design-patterns-and-principles)
  - [Design Patterns](#design-patterns)
  - [Architectural Principles](#architectural-principles)
- [Running Tests](#running-tests)

## Prerequisites

- Go 1.23 or higher
- Docker
- Docker Compose

## Getting Started

### Clone the Repository

```bash
git clone https://github.com/yourusername/rest_api_go.git
cd rest_api_go
```

### Environment Setup

Set up your environment variables for database connection if you're not using docker-compose:

```bash
export POSTGRES_HOST=localhost
export POSTGRES_PORT=5432
export POSTGRES_USER=yourusername
export POSTGRES_PASSWORD=yourpassword
export POSTGRES_DB=yourdbname
export PORT=8080
```

### Build and Run with Docker

1. **Build Docker Images**

   Ensure Docker is running, then build the images:

   ```bash
   docker-compose up --build
   ```

2. **Access the Application**

   - **REST API**: Accessible via `http://localhost:8080`
   - **GraphQL Playground**: Accessible via `http://localhost:8080/graphql`

## API Endpoints 
#### Either use CURL or Postman
### Authentication

- **Login**

  ```http
  POST /login
  ```

  **Payload:**

  ```json
  {
    "username": "username",
    "password": "password"
  }
  ```

### Devices

- **Get All Devices**

  ```http
  GET /devices
  ```

- **Get Device by ID**

  ```http
  GET /devices/{id}
  ```

- **Create Device**

  ```http
  POST /devices
  ```

  **Payload:**

  ```json
  {
    "name": "Device Name",
    "version": "1.0"
  }
  ```

### GraphQL Query Example

```graphql
query {
  devices {
    id
    name
    version
  }
  
  device(id: "1") {
    id
    name
    version
  }
}
```

## Design Patterns and Principles

### Design Patterns

- **Repository Pattern**: This pattern separates data access logic from business logic, making it easier to test and maintain. It abstracts the database operations, enabling mock implementations for testing.

- **Service Layer**: Provides a higher level of abstraction over the repository, implementing business logic and interacting with multiple repositories if needed.

- **Middleware Pattern**: Used for cross-cutting concerns such as authentication and logging. The JWT middleware implements token validation for protected routes.

### Architectural Principles

- **REST (Representational State Transfer)**: The API follows REST principles with stateless operations, resource-based URLs, and standard HTTP methods (GET, POST).

- **HATEOAS (Hypermedia as the Engine of Application State)**: While not fully implemented, the API can be extended to include links in responses to enable clients to navigate the API dynamically.

- **GraphQL**: Provides flexibility over traditional REST by allowing clients to request exactly the data they need, reducing over-fetching and under-fetching issues.

## Running Tests ( Not implemented )

To run unit tests, execute:

```bash
go test ./...
```
