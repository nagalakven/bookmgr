
# CLI User Guide: Book Management System

The `bookmgr` CLI provides commands to interact with the book management system.
The available commands, their options, sample responses and examples on how to use them are outlined below.

---

## Table of Contents
1. [Add a New Book](#add-a-new-book)
2. [Delete a Book](#delete-a-book)
3. [Delete Books by Filter](#delete-books-by-filter)
4. [Update a Book](#update-a-book)
5. [Create a Collection](#create-a-collection)
6. [Add Books to Collection](#add-books-to-collection)
7. [Remove Books from Collection](#remove-books-from-collection)
8. [Delete a Collection](#delete-a-collection)
9. [List All Books](#list-all-books)
10. [Filter Books](#filter-books)
11. [List All Collections](#list-all-collections)
12. [List All Books in a Collection](#list-all-books-in-a-collection)

--- 

## Commands

### Add a New Book
Add a new book to the system by providing the required and optional details.

```
bookmgr add-book [OPTIONS]
```

**Options:**

- `-t or --title <title>`
  *reqired.* Title of the book.
  Example: `--title "Book Title"`

- `-a or --author`
  *required.* Author of the book.
  Example: `--author "Author Name"`

- `-p or --published-date`
  *required.* Publication date of the book (format: YYYY-MM-DD).
  Example: `--published-date "2025-01-01"`

- `-b or --published-by`
  *optional.* Name of the publisher of the book.
  Example: `--published-by "Publisher Name"`

- `-c or --co-authors <names>`
  *optional.* A comma-separated list of co-authors of the book, if any.
  Example: `--co-authors "Co-Author 1, Co-Author 2"`

- `-g or --genre`
  *optional.* Genre of the book.
  Example: `--genre "Fiction"`

- `-d or --description`
  *optional.* Short description of the book.
  Example: `--description "A fictional story"`

- `-e or --edition`
  *optional.* Edition of the book.
  Example: `--edition "1"`

- `-r or --recommended-age`
  *optional.* The recommended age range for the book.
  Example: `--recommended-age "10+"`

- `-s or --isbn`
  *optional.* International Standar Book Number of the book.
  Example: `--isbn "978-1-945209-05-5"`

#### Example Commands and Responses:

##### Example-1: Add a New Book with Required Fields

```
bookmgr add-book --title "Book Title 1" --author "Author Name" --published-date "2025-01-01"
```

**Response:**

```
Book "Book Title 1" added successfully!
ID: 12345
``` 

##### Example-2: Add a New Book with All Fields

```
bookmgr add-book --title "Book Title 2" --author "Author Name" --published-date "2025-01-01" --published-by "Publisher Name" --co-authors "Co-Author 1, Co-Author 2" --genre "Fiction" --description "A fictional story" --edition "1" --recommended-age "10+" --isbn "978-1-945209-05-5"
```

**Response:**

```
Book "Book Title 2" added successfully!
ID: 23423
```

### Delete a Book
Remove a specific book from the system using its unique identifier.

```
bookmgr delete-book [OPTIONS]
```

**Options:**

- `-i or --id <id>`
  *required.* The unique identifier of the book to be deleted.
  Example: `--id 12345`

#### Example Commands and Responses:

##### Example-1: Delete a Book by ID

```
bookmgr delete-book --id 12345
```

**Response:**

```
Book with ID 12345 deleted successfully!
```

### Delete Books by Filter
Remove books from the system based on specific filter criteria, such as author, genre, or a range of publication dates.

```
bookmgr delete-books [OPTIONS]
```

**Options:**

- `-a or --author`
  *optional.* Delete books written by the specified author.
  Example: `--author "Author Name"`

- `-g or --genre`
  *optional.* Delete books belonging to the speficied genre.
  Exampl: `--genre "Fiction"`

- `-p or --published-date`
  *optional.* Delete books published on a specific date (format: `YYYY-MM-DD`) or within an inclusive range (format: `YYYY-MM-DD` to `YYYY-MM-DD`).
  Example: `--published-date "2025-01-01"`
  Example: `--published-date "2025-01-01 to 2025-02-01"`

#### Example Commands and Responses:

##### Example-1: Delete Books by Author

```
bookmgr delete-books --author "Author Name"
```

**Response:**

```
Books by "Author Name" deleted successfully!
```

##### Example-2: Delete Books by Publication Date Range

```
bookmgr delete-books --published-date "2025-01-01 to 2025-02-01"
```

**Response:**

```
Books published between "2025-01-01" and "2025-02-01" deleted successfully!
```

### Update a Book
Modify the details of an existing book in the system.

```
bookmgr update-book [OPTIONS]
```

**Options:**

- `-i or --id <id>`
  *required.* The unique identifier of the book to be updated.
  Example: `--id 12345`

- `-t or --title <title>`
  *optional.* The updated title of the book.
  Example: `--title "Updated Title"`

- `-a or --author <author>`
  *optional.* The updated author of the book.
  Example: `--author "Updated Author"` 

- `-p or --published-date <YYYY-MM-DD>`
  *optional.* The updated publication date of the book.
  Example: `--published-date "2025-01-01"`

- `-b or --published-by <publisher>`
  *optional.* The updated publisher of the book.
  Example: `--published-by "Updated Publisher"`

- `-c or --co-authors <co-authors>`
  *optional.* A comma-separated list of updated co-authors of the book.
  Example: `--co-authors "Co-Author 1, Co-Author 2"`

- `-g or --genre <genre>`
  *optional.* The updated genre of the book.
  Example: `--genre "Non-fiction"`

- `-d or --description <description>`
  *optional.* The updated description of the book.
  Example: `--description "An updated description of the book"`

- `-e or --edition <edition>`
  *optional.* The updated edition of the book.
  Example: `--edition "2"`

- `-r or --recommended-age <age-range>`
  *optional.* The updated recommended age range for the book.
  Example: `--recommended-age "10+"

- `-s or --isbn <isbn>`
  *optional.* The updated ISBN number of the book.
  Example: `--isbn "978-1-123456-78-9"`

#### Example Commands and Responses:

##### Example-1: Update Book Title and Author

```
bookmgr update-book --id 12345 --title "Updated Title" --author "Updated Author"
```

**Response:**

##### Example-2: Update Book Genre and Publication Date

```
bookmgr update-book --id 12345 --genre "Non-fiction" --published-date "2022-01-01"
```

**Response:**

```
Book ID 12345 updated successfully!
```

### Create a Collection
Create a new collection for books.

```
bookmgr create-collection --name "New Collection"
```

**Options:**
- `-n or --name <collection-name>`
  *required.* The name of the new collection.
  Example: `--name "My Favorites"`

#### Example Commands and Responses:

##### Example-1: Add a New Collection

```
bookmgr create-collection --name "New Collection"
```

**Response:**

```
Collection "New Collection" created successfully!
ID: 67890
```

### Add Books to Collection
Add one or more books to an existing collection.

```
bookmgr add-books-to-collection [OPTIONS]
```

**Options:**

- `-c or --collection-id <collection-id>`
  *required.* The unique identifier of the collection to which books should be added.
  Example: `--collection-id 67890`

- `-b or --book-ids <book-ids>`
  *required.* A comma-separated list of book identifiers to add to the collection.
  Example: `--book-ids "12345, 67891"`

#### Example Commands and Responses:

##### Example-1: Add Books to a Collection

```
bookmgr add-books-to-collection --collection-id 67890 --book-ids 12345,67891
```

**Response:**

```
Books added to collection ID 67890 successfully!
```

### Remove Books from Collection
Remove one or more books from an existing collection.

```
bookmgr remove-books-from-collection [OPTIONS]
```

**Options:**

- `-c, --collection-id <collection-id>`
  *required.* The unique identifier of the collection from which books should be removed.
  Example: `--collection-id 67890`

- `-b or --book-ids <book-ids>`
  *required.* A comma-separated list of book identifiers to remove from the collection.
  Example: `--book-ids "12345, 67891"`

#### Example Commands and Responses:

##### Example-1: Remove Books from a Collection

```
bookmgr remove-books-from-collection --collection-id 67890 --book-ids 12345,67891
```

**Response:**

```
Books removed from collection ID 67890 successfully!
```

### Delete a Collection
Delete a collection from the system.

```
bookmgr delete-collection [OPTIONS]
```

**Options:**

- `-i or --id <collection-id>`
  *required.* The unique identifier of the collection to be deleted.
  Example: `--id 67890`

#### Example Commands and Responses:

##### Example-1: Delete a Collection

```
bookmgr delete-collection --id 67890
```

**Response:**

```
Collection ID 67890 deleted successfully!
```

### List All Books
Retrieve and display all books stored in the book management system.

```
bookmgr list-books [OPTIONS]
```

**Options:**

- `-l or --limit <number>`
  *optional.* The maximum number of books to display.
  Example: `--limit 10`

- `-s, --sort-by <field>`
  *optional.* Sort the list of books by a specific field (e.g., `title`, `author`, `published-date`).
  Example: `--sort-by "title"`

- `-o, --order <asc|desc>`  
  *optional.* Specify the sorting order: ascending (`asc`) or descending (`desc`).  
  Example: `--order "asc"`

- `-f, --filter <criteria>`  
  *optional.* Apply a filter to list books based on specific criteria (e.g., `--filter "genre=Fiction"` or `--filter "author=Author Name"`).  
  Example: `--filter "genre=Non-fiction"`

#### Example Commands and Responses:

##### Example-1: List All Books

```
bookmgr list-books
```

**Response:**

```
ID      Title               Author        Genre         Published Date
------------------------------------------------------------------------
67890   Book Title 2        Author 2      Non-fiction   2023-07-15
12345   Book Title 1        Author 1      Fiction       2025-01-01
```

##### Example-2: List Books Sorted by Title

```
bookmgr list-books --sort-by "title" --order "asc"
```

**Response:**

```
ID      Title               Author        Genre         Published Date
------------------------------------------------------------------------
12345   Book Title 1        Author 1      Fiction       2025-01-01
67890   Book Title 2        Author 2      Non-fiction   2023-07-15
```

##### Example-3: List Books with Filters

```
bookmgr list-books --filter "author=Author 1"
```

**Response:**

```
ID      Title               Author        Genre         Published Date
------------------------------------------------------------------------
12345   Book Title 1        Author 1      Fiction       2025-01-01
```

### Filter Books
Filter books based on the specified options.

```
bookmgr filter-books [OPTIONS]
```

**Options:**
- `-a, --author <name>`
  *optional.* Filter the list of books by the specified author's name.
  Example: `--author "Author Name"`

- `-g, --genre <genre>`
  *optional.* Filter the list of books by the specified genre.
  Example: `--genre "Fiction"`

- `-s, --start-date <YYYY-MM-DD>`
  *optional.* Filter books on or after the specified start date.
  Example: `--start-date "2025-01-01"`

- `-e, --end-date <YYYY-MM-DD>`
  *optional.* Filter books on or before the specified end date.
  Example: `--end-date "2025-01-01"`
  Note: If both --start-date and --end-date are provided, the command will filter books published between those dates (inclusive).

- `-o, --output <format>`
  *optional.* Defines the format of the output.
  Values:
  - `full`: Displays all configured details for each book (default).
  - `short`: Displays only the titles of the books.
  - `json`: Displays all configured details for each book in JSON format.
  - `csv`: Displays all configured details for each book in CSV format.
  - `custom=<fields>`: Customizes the display to include only specific fields for each book.
  Example: `--output "custom=title,genre"`
  Available Fields: title, author, genre, description, published_date, recommended_age.

- `-l, --limit <number>`
  *optional.* Limits the number of results displayed. Defaults to all if not specified.
  Example: `--limit 10`


#### Example Commands and Responses:

##### Example-1: Filter by Author
```
bookmgr filter-books -a "Author Name"
```

**Response:**
```
Found 2 books by "Author Name":

1. ID: 3456 - Title: "Book Title 1" - Published: 2023-01-01 - Genre: Fiction - Description: A Fiction Story
2. ID: 7890 - Title: "Book Title 2" - Published: 2022-06-15 - Genre: Mystery - Description: A Mystery Novel - Recommended-Age: 18+
```

##### Example-2: Filter by start date

```
bookmgr filter-books --start-date "2023-01-01"
```

**Response:**
```
Found 1 book published on 2023-01-01:

1. ID: 3456 - Title: "Book Title 1" - Genre: Fiction
```

##### Example-3: Filter by range of dates

```
bookmgr filter-books -s "2022-01-01" -e "2023-01-01"
```

**Response:**
```
Found 2 books published between 2022-01-01 and 2023-01-01:

1. ID: 7890 - Title: "Book Title 2" - Published: 2022-06-15 - Genre: Mystery - Description: A Mystery Novel - Recommended-Age: 18+
2. ID: 3456 - Title: "Book Title 1" - Published: 2023-01-01 - Genre: Fiction - Description: A Fiction Story
```

### List All Collections
List all collections created in the system.

```
bookmgr list-collections
```

#### Example Command and Response:

##### Example-1: List All Collections

```
bookmgr list-collections
```

**Response:**

```
ID      Name           
---------------------
67890  New Collection
67891  My Favorites  
```

### List All Books in a Collection
List all books that belong to the specified collection in the system.

```
bookmgr list-books-in-collection [OPTIONS]
```

**Options**:

- `-n or --name <collection-name>`
  *optional.* The name of the collection to list books from.
  Example: `--name "My Favorites"`

- `-i or --id <collection-id>`
  *optional.* The unique identifier of the collection to filter by.
  Example: `--id 67890`

#### Example Command and Response:

##### Example-1: List All Books in a Collection by Collection Name

```
bookmgr list-books-in-collection --name "My Favorites"
```

**Response:**

ID      Title                     Author        Genre         Published Date
-----------------------------------------------------------------------------
12345   Book Title 1              Author 1      Fiction       2025-01-01
56478   Another Fiction Book      Author 3      Fiction       2023-07-15


##### Example-2: List All Books in a Collection by unique collection identifier

```
bookmgr list-books-in-collection --id 67890
```

**Response:**

ID      Title                     Author        Genre         Published Date
-----------------------------------------------------------------------------
12345   Book Title 1              Author 1      Fiction       2025-01-01
56478   Another Fiction Book      Author 3      Fiction       2023-07-15


