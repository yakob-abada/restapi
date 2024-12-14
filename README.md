# GO Rest API

This project demonstrates the use of the GIN framework in Go, following the Domain-Driven Design (DDD) architecture pattern. It integrates MariaDB as the database to persist domain data, ensuring scalability, maintainability, and separation of concerns.

### Getting Started with the Project

This project is containerized using **Docker Compose** and includes both a **MariaDB** container and a **phpMyAdmin** container for easy database management. Follow the steps below to set up the environment and start interacting with the API.

#### 1. **Start the Containers**
To get the MariaDB and phpMyAdmin containers up and running, use **Docker Compose**:

```bash
docker-compose up
```

This will:
- Launch a **MariaDB** container for the database.
- Launch a **phpMyAdmin** container for convenient database management via a web interface.

#### 2. **Set Up the Database**
After the containers are up, run the following command to set up the database schema and populate it with some dummy data:

```bash
make db_setup
```

This command will:
- Create the necessary tables in the MariaDB database.
- Insert some dummy data for testing purposes.

#### 3. **Run the Application**
Now, you're ready to start the application. Run the following command to launch the server:

```bash
make run
```

This will start the web server and make the application available for use.

---

### Available Endpoints

Once the application is running, you can access the following API endpoints to interact with the system:

1. **Explore Matched Profiles**  
   URL: `http://localhost:8080/profile_explore`  
   **Description**: Fetches profiles that are matched with the currently logged-in user. This could return a list of recommended profiles based on preferences, interests, etc.

2. **View Matches**  
   URL: `http://localhost:8080/match`  
   **Description**: Returns a list of users who are matched with the currently logged-in user. This endpoint provides details of users with whom a match has been established.

3. **View Likes**  
   URL: `http://localhost:8080/like`  
   **Description**: Returns a list of users who have liked the currently logged-in user. This can help the user see who has expressed interest in them.

---

Here's an improved and more detailed version of the **Summary of Commands** section. It includes clear descriptions of each command and makes it easier for users to understand the purpose of each action:

---

### Summary of Commands

Below is a list of key commands to set up, test, and run the application, along with brief descriptions of what each command does:

- **Start Containers**  
  Command: `docker-compose up`  
  **Description**: Spins up the Docker containers for the application, including **MariaDB** (database) and **phpMyAdmin** (for database management).

- **Set Up Database**  
  Command: `make db_setup`  
  **Description**: Runs the necessary database setup, which typically involves configuring the initial state of the database (e.g., creating basic structures, setting up environment variables).

- **Create Tables in Database**  
  Command: `make migration_up`  
  **Description**: Applies database migrations to create the required tables and schemas in **MariaDB**. This is useful when setting up the database schema or when migrating to a new version.

- **Drop Tables in Database**  
  Command: `make migration_down`  
  **Description**: Rolls back any applied migrations and drops the tables from the database. This can be useful when resetting the database or testing with a clean slate.

- **Add Dummy Data to Tables**  
  Command: `make fixtures`  
  **Description**: Populates the database with dummy or test data. This is useful for quickly setting up realistic data for testing or development.

- **Run Unit Tests**  
  Command: `make unit_test`  
  **Description**: Executes the unit tests for the application, ensuring that the code works as expected. Useful for verifying the correctness of individual components.

- **Run the Application**  
  Command: `make run`  
  **Description**: Starts the application server, making the API available to interact with. Typically runs the server on `localhost:8080` (or a custom port depending on your configuration).

---

### Available Endpoints

Once the application is running, you can test and interact with the following API endpoints:

1. **Explore Matched Profiles**  
   URL: `http://localhost:8080/profile_explore`  
   **Description**: Fetches profiles that are matched with the logged-in user.

2. **View Matches**  
   URL: `http://localhost:8080/match`  
   **Description**: Returns a list of users who are matched with the logged-in user.

3. **View Likes**  
   URL: `http://localhost:8080/like`  
   **Description**: Returns a list of users who have liked the logged-in user.
