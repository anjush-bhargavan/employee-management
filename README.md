# Employee Management System API

This project is a simple Employee Management System API built with Golang. The API allows for the creation, retrieval, updating, and deletion of employee records. It integrates with PostgreSQL for data storage, Echo for the web framework, Swagger for API documentation, and Redis for caching.

## Technologies Used

- **Golang**: Main programming language.
- **PostgreSQL**: Database for storing employee records.
- **Echo**: Web framework for handling HTTP requests.
- **Swagger**: For API documentation and testing.
- **Redis**: For caching purposes to improve read performance.

## Getting Started

### Prerequisites

- Golang (version 1.16 or higher)
- PostgreSQL
- Redis
- Docker (optional, for running the application in a container)

### Setup Instructions

1. **Clone the repository**

   ```bash
   git clone https://github.com/anjush-bhargavan/employee-management.git
   cd employee-management

2. Set up Redis & PostgreSQL

    Make sure Redis is running on its default port (6379).

3. Create .env file

    ```bash
      DBHOST="localhost"
      DBUSER="postgres"
      DBPASSWORD="#######"
      DBNAME="employee_management"
      DBPORT="5432"
      DBSSLMODE="disable"
      
      REDISHOST="localhost:6379"

4. Run the application
      
      ```bash
      go mod tidy
      go run main.go

5. Run with Docker (optional)

    If you prefer to run the application using Docker, you can use the provided Dockerfile.

      ```bash
      docker build -t employee-management .
      docker run -p 8080:8080 --env-file .env employee-management
6. API Documentation

    The API documentation is accessible via Swagger UI. Once the application is running, navigate to the following URL in your web browser:

      ```bash
      http://localhost:8080/swagger/index.html
