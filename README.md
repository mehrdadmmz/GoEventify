# GoEventify 

A lightweight **RESTful API** written in Go for creating and managing events. Built with the [Gin Web Framework](https://github.com/gin-gonic/gin) and a SQLite database, this project supports:

- **User registration & login**  
- **JWT-based authentication**  
- **Event creation, updating, and deletion**  
- **User registration for events (and cancellation)**  

This makes it a handy starter template or reference for a small event management system in Go.

---

## Table of Contents

1. [Features](#features)  
2. [Tech Stack](#tech-stack)  
3. [Project Structure](#project-structure)  
4. [Installation](#installation)  
5. [Usage](#usage)  
6. [API Endpoints](#api-endpoints)  
7. [License](#license)  
8. [Contact](#contact)  

---

## Features

1. **User Signup / Login**  
   - Secure password hashing (using `bcrypt`).  
   - JWT token generation for authentication.  

2. **JWT-Based Authentication**  
   - Middleware checks for a valid token in the `Authorization` header.  
   - Only authenticated users can create, update, or delete events.  

3. **Event Management**  
   - Create, read, update, and delete events.  
   - Only the event owner can update or delete their event.  

4. **Event Registration**  
   - Authenticated users can register for events.  
   - Users can also cancel their registration.  

5. **Database**  
   - SQLite database with three tables (`users`, `events`, `registrations`).  

---

## Tech Stack

- **Go** (1.x and above)  
- **Gin** (for HTTP routing and middleware)  
- **SQLite** (local file-based database)  
- **JWT** (JSON Web Token for secure authentication)  
- **Bcrypt** (for secure password hashing)  

---

## Project Structure
```bash
REST_PROJECT/
│
├── db/
│   └── db.go                # Database initialization & table creation
│
├── middlewares/
│   └── auth.go              # JWT authentication middleware
│
├── models/
│   ├── user.go              # User model and methods
│   ├── event.go             # Event model and methods
│   └── register.go          # Event registration logic
│
├── routs/
│   ├── routs.go             # Main router registration
│   ├── users.go             # Handlers for user signup, login
│   ├── events.go            # Handlers for event CRUD
│   └── registers.go         # Handlers for event registrations
│
├── utils/
│   ├── hash.go              # Password hashing & verification
│   └── token.go             # (Example) JWT generation & verification
│
└── main.go                  # Application entry point
```


## Installation

1. **Prerequisites**  
   - [Go](https://go.dev/dl/) (version 1.18+ recommended)  
   - Git (to clone this repository)  

3. **Clone the repository**  
   ```bash
   git clone https://github.com/your-username/GoEventify.git
   cd GoEventify

4. **Initialize / Vendor Dependencies**
   ```bash
   go mod tidy
   ```
   This ensures all necessary Go modules are downloaded.

5. **Database**
   - By default, the project uses a local api.db file. The code in db/db.go will create it automatically if it does not exist.
   - No extra setup required for SQLite.

## Usage

1. **Start the Server**
   ```bash
   go run main.go
   ```
   The application typically runs on port :8080

3. **Test the Endpoints**
   - Use an API client like Postman or cURL to interact with the API.
   - Example:
     ```bash
     curl --location --request GET 'http://localhost:8080/events'
      ```
     
5. **Authentication**
   - Create a user via POST /signup (returns success if valid).
   - Log in via POST /login (you’ll receive a JWT in the response).
   - For protected endpoints, include the JWT as a Bearer token in the Authorization header:
     ```bash
     Authorization: <JWT_TOKEN>
     ```

## API Endpoints

| **Method** | **Endpoint**               | **Description**                                   | **Auth Required** |
|------------|----------------------------|---------------------------------------------------|-------------------|
| **POST**   | `/signup`                 | Register a new user                              | No                |
| **POST**   | `/login`                  | Log in an existing user, returns JWT token       | No                |
| **GET**    | `/events`                 | Get a list of all events                         | No                |
| **GET**    | `/events/:id`             | Get details of a specific event                  | No                |
| **POST**   | `/events`                 | Create a new event                               | Yes (token)       |
| **PUT**    | `/events/:id`             | Update an existing event                         | Yes (token)       |
| **DELETE** | `/events/:id`             | Delete an existing event                         | Yes (token)       |
| **POST**   | `/events/:id/register`    | Register the user for an event                   | Yes (token)       |
| **DELETE** | `/events/:id/register`    | Cancel the user’s registration for an event      | Yes (token)       |

## Request & Response Examples

1. **Signup**
   - Request:
     ```bash
     POST /signup
      {
        "email": "newuser@example.com",
        "password": "mypassword"
      }
     ```
   - Response:
     ```json
     {
      "message": "User created successfully!"
     }
     ```
2. **Login**
   - Request:
     ```bash
     POST /login
      {
        "email": "newuser@example.com",
        "password": "mypassword"
      }
     ```
   - Response:
     ```json
     {
      "message": "Login successful",
      "token": "<your-jwt-token-here>"
     }
     ```
3. **Create Event**
   - Request (requires Authorization: <JWT_TOKEN> header):
     ```bash
     POST /events
      {
        "name": "Tech Conference",
        "description": "All about GoLang",
        "location": "Virtual",
        "date_time": "2025-03-15T10:00:00Z"
      }
     ```
   - Response:
     ```json
     {
        "message": "Event created!",
        "event": {
          "id": 1,
          "name": "Tech Conference",
          "description": "All about GoLang",
          "location": "Virtual",
          "date_time": "2025-03-15T10:00:00Z",
          "user_id": 5
        }
      }
     ```

## License
```nginx
MIT License
```


     
