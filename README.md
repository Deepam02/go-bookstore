# Go Bookstore Management System

A **backend-only** bookstore management system built with **Golang** and **MySQL**. This project focuses on handling book-related operations using RESTful APIs. It uses **GORM** for ORM and **gorilla/mux** for routing.

## Features

- **GORM (ORM)**: Efficiently interacts with a MySQL database to handle book records.
- **gorilla/mux**: Powerful router for API endpoints.
- **JSON Marshal/Unmarshal**: Processes request and response data in JSON format.
- **Database Integration**: Connects to a **MySQL** database (supports both local and online setups).

## API Endpoints

This project exposes the following RESTful APIs to manage books:

1. **GET /book**
   - **Description**: Retrieves all books.
   - **Response**: JSON array of all books in the database.

2. **GET /book/{bookId}**
   - **Description**: Retrieves a single book by its unique ID.
   - **Response**: JSON object with book details.

3. **POST /book**
   - **Description**: Creates a new book record.
   - **Request Body**: JSON object with book details.
   - **Response**: JSON object with the newly created book.

4. **DELETE /book/{bookId}**
   - **Description**: Deletes a book by its ID.
   - **Response**: JSON object with a success message.

5. **PUT /book/{bookId}**
   - **Description**: Updates an existing book by its ID.
   - **Request Body**: JSON object with updated book details.
   - **Response**: JSON object with the updated book.

## Database Integration

- The system is connected to a **MySQL** database. While this implementation uses an online database, you can configure the connection for either a **local** or **cloud-hosted** MySQL instance.
  
- Database interactions are handled using **GORM**, which abstracts away raw SQL queries and allows seamless interactions with the database.

## Technologies Used

- **Go (Golang)**: The core language for backend development.
- **GORM**: Object Relational Mapping (ORM) library for Go.
- **gorilla/mux**: A powerful routing library for Go to handle HTTP requests.
- **MySQL**: Relational database for storing book data.

## How it Works

1. **Routing**: API routes are defined using the **gorilla/mux** router, which maps incoming HTTP requests to the appropriate handler functions.
2. **Database Operations**: When a request is made, the handler interacts with the MySQL database through **GORM** to perform CRUD operations (Create, Read, Update, Delete).
3. **Request/Response Handling**: All data is exchanged in **JSON** format. Incoming requests are unmarshalled into Go structs, and responses are marshalled into JSON before sending them back.



### Contact

For questions or suggestions, feel free to open an issue or contact me directly!

