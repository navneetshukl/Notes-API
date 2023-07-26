# JWT Authentication with Golang, Gin, and Gorm

This is a simple project that demonstrates how to implement JWT (JSON Web Token) based authentication in a Golang web application using the Gin web framework and Gorm as the ORM (Object-Relational Mapping) library for PostgreSQL database.

## Table of Contents

- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Project Structure](#project-structure)
- [Endpoints](#endpoints)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Getting Started

### Prerequisites

Before running this project, you need to have the following software installed:

- Golang (version 1.15 or higher)
- Docker
- Docker Compose

### Installation

1.  Clone the repository

        git clone https://github.com/your-username/your-repo-name.git
        cd your-repo-name

2. Create the necessary environment files:

    Create a `.env` file in the root of the project and add the following environment variables:

         POSTGRES_USER=your_postgres_username
 
         POSTGRES_PASSWORD=your_postgres_password
 
         POSTGRES_DB=JWT_Auth
   
         SECRET=your_jwt_secret_key



3. Build and run the project using Docker Compose:

          docker-compose up --build



The application will be accessible at `http://localhost:8080`.

## Project Structure

The project follows a modular structure to keep the code organized and maintainable. Here's a brief overview of the project structure:

- `controllers/`: Contains the Gin controller functions responsible for handling HTTP requests and responses.
- `database/`: Contains the database setup and connection code using Gorm.
- `models/`: Contains the Gorm model structs representing database tables.
- `main.go`: The entry point of the application.

## Endpoints

- `POST /signup`: Sign up a new user with a name, email, and password. Returns a success message if the user is created successfully.
- `POST /login`: Authenticate a user with email and password, and issue a JWT token on successful login.
- `GET /notes`: Get all notes associated with the authenticated user.
- `POST /notes`: Create a new note for the authenticated user.
- `GET /notes/:title`: Get a single note by title for the authenticated user.
- `PUT /notes/:title`: Update a note's description by title for the authenticated user.
- `DELETE /notes`: Delete all notes associated with the authenticated user.
- `DELETE /notes/:title`: Delete a single note by title for the authenticated user.

## Usage

1. Sign up a new user using the `POST /signup` endpoint with a JSON request body containing `name`, `email`, and `password`.

2. Log in with the newly created user using the `POST /login` endpoint with a JSON request body containing `email` and `password`. The endpoint will return a JWT token that you need to include in subsequent requests as a bearer token in the `Authorization` header.

3. Use the provided token to access the protected endpoints (`/notes`, `/notes/:title`, etc.). Include the token in the `Authorization` header with the value `Bearer <token>`.

4. Use the other endpoints (`/notes`, `/notes/:title`, etc.) to manage notes for the authenticated user.

## Contributing

Contributions to this project are welcome. If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.


