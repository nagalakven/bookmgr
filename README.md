# bookmgr
This is a simple **Book Management Software** which provides REST APIs and CLI interraction for personal/small business library management.
This system allows the users to:
* Add and manage books with basic information [title, author, publication date, edition, description, genre, reader's suggested age]
* Create a manage user-defined collections of books
* List books, collections and filter lists by author, genre or a range of publication dates

## Features
* **REST APIs**: A set of RESTful endpoints are exposed by the system to interract with books and collections.
* **Relational Database**: The data [books,collections] are stored in a relational database like Sqlite.
                           For the purpose of this assignment, in-memory is used.
* **CLI**: A command-line interface is provided for alternate means of user interraction with the data.
* **Dockerized Deployment**: Application is packaged using Docker for easy deployment and scaling.

## Technologies
* **Backend**: Go,
* **Database**: In-memory (any relational database)
* **CLI**: Go, Linux-syntax
* **Version Control**: Git

## Setup
1. Clone repository:

   ```bash
   git clone https://github.com/nagalakven/bookmgr.git
   ```

2. Navigate to project directory:

   ```bash
   cd bookmgr
   ```

## Documentations

The design documents for the assignment are found in the repository under bookmgr/docs path.

1. [CLI user specification](docs/cli_user_guide.md)
2. [Database design specification](docs/database_design.md)
3. [REST API documentation](docs/rest-api.md)
4. [OpenAPI specification](docs/rest-api.yaml)
5. [Unit test results](docs/tests/test_results.md)
6. [Sample log file](log/app.20250128_223926.log)

## Assumptions and Future Enhancements

**Assumptions**

The following assumptions are made for the purpose of this assignment.
This is a minimalistic implementation and needs more optimization.

1. Instead of integrating with a database, the service utilizes an in-memory storage.
2. Code is structured to flow as follows: main -> handler -> service -> respository.
3. The server is exposed via simple HTTP endpoints.
4. Basic validations are handled.
5. The service handles CRUD operations for books. These are the operations implemented:
   - Add a book
   - Update a book
   - Delete a book
   - List all books
6. Basic logging is done. Logs can be found in bookmgr/logs path.

**Enhancements**

The following as some of the future enhancements for the service.

- API authentication using Bearer token, OAuth2
- Batch processing multiple requests
- Database integration
- Maintaing context flow

