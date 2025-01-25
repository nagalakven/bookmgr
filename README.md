# bookmgr
This is a simple **Book Management Software** which provides REST APIs and CLI interraction for personal/small business library management.
This system allows the users to:
* Add and manage books with basic information [title, author, publication date, edition, description, genre, reader's suggested age]
* Create a manage user-defined collections of books
* List books, collections and filter lists by author, genre or a range of publication dates

## Features
* **REST APIs**: A set of RESTful endpoints are exposed by the system to interract with books and collections.
* **Relational Database**: The data [books,collections] are stored in a relational database like Sqlite.
* **CLI**: A command-line interface is provided for alternate means of user interraction with the data.
* **Dockerized Deployment**: Application is packaged using Docker for easy deployment and scaling.

## Technologies
* **Backend**: Go, Flask
* **Database**: Sqlite
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

[TBU further]
   
