# Bank API Project

This project implements a simple bank API that allows customers to log in, make payments to merchants, and log out. It uses clean architecture principles and JWT for authentication.

## Features

- Customer login with JWT token generation
- Secure payment processing between customers and merchants
- Logout functionality
- Activity logging for all operations
- JSON file-based data storage (for demonstration purposes)

## Project Structure

```
project/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── api/
│   │   ├── handlers/           # HTTP request handlers
│   │   └── routes.go           # API route definitions
│   ├── domain/                 # Domain models
│   ├── repository/             # Data access layer
│   ├── usecase/                # Business logic
│   └── middleware/             # HTTP middlewares
├── data/                       # JSON data files
├── go.mod
└── README.md
```

## Prerequisites

- Go 1.16 or later
- Git

## Getting Started

1. Clone the repository:

   ```
   git clone https://github.com/onuda22/simple_bank.git
   cd bank-api-project
   ```

2. Install dependencies:

   ```
   go mod tidy
   ```

3. Set up the JWT secret on .env file:

   ```
   JWT_SECRET=your_secret_key_here
   ```

4. Run the application:
   ```
   go run cmd/main.go
   ```

The server will start on `http://localhost:8080`.

## API Endpoints

### Login

- **URL**: `/login`
- **Method**: `POST`
- **Body**:
  ```json
  {
    "username": "john_doe",
    "password": "securePassword123"
  }
  ```
- **Response**:
  ```json
  {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
  ```

### Make Payment

- **URL**: `/payment`
- **Method**: `POST`
- **Headers**:
  - `Authorization: Bearer <your_jwt_token>`
- **Body**:
  ```json
  {
    "merchant_id": "M001",
    "amount": 10500
  }
  ```
- **Response**:
  ```json
  {
    "message": "Payment successful"
  }
  ```

### Logout

- **URL**: `/logout`
- **Method**: `POST`
- **Headers**:
  - `Authorization: Bearer <your_jwt_token>`
- **Response**:
  ```json
  {
    "message": "Logout successful"
  }
  ```
