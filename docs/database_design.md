# Relation Database Tables for Book Management System

This file briefs about various tables, their coulmns, column types, relations, constraints and index for the Book Management System.

## Tables

### Books Table:

**Column-1: ID**

```
Name: id
Data Type: UUID
Description: Unique identifier for the book.
Constraints: Primary Key, auto-generated
Index: Primary Key
```

**Column-2: Title**

```
Name: title
Data Type: VARCHAR(255)
Description: Title of the book (required).
Constraints: NOT NULL
```

**Column-3: Author**

```
Name: author
Data Type: VARCHAR(255)
Description: Author of the book (required).
Constraints: NOT NULL
```

**Column-4: Published Date**

```
Name: published_date
Data Type: DATE
Description: Publication date of the book in YYYY-MM-DD format (required).
Constraints: NOT NULL
```

**Column-5: Published By**

```
Name: published_by
Data Type: VARCHAR(255)
Description: Name of the publisher of the book (optional).
```

**Column-6: Co-authors**

```
Name: co_authors
Data Type: JSON
Description: JSON array of co_authors (optional).
```

**Column-7: Genre**

```
Name: genre
Data Type: VARCHAR(100)
Description: Genre of the book (optional).
```

**Column-8: Description**

```
Name: description
Data Type: TEXT
Description: A short description of the book (optional).
```

**Column-9: Edition**

```
Name: edition
Data Type: VARCHAR(50)
Description: Edition of the book (optional).
```

**Column-10: Recommended Age**

```
Name: recommended_age
Data Type: VARCHAR(50)
Description: Recommended target age range for the book (optional).
```

**Column-11: ISBN**

```
Name: isbn
Data Type: VARCHAR(20)
Description: ISBN of the book (optional).
Constraints: UNIQUE
Index: Unique Index
```

### Collections Table:

**Column-1: ID**

```
Name: id
Data Type: UUID
Description: Unique identifier for the collection.
Constraints: Primary key, auto-generated 
Index: Primary Key
```

**Column-2: Name**

```
Name: name
Data Type: VARCHAR(255)
Description: Name of the collection (required).
Constraints: NOT NULL, UNIQUE
Index: Unique Index
```

### Books_Collections Table:

This is the *join table*. It holds many-to-many relationship between Books and Collections.

**Column-1: ID**

```
Name: id
Data Type: UUID
Description: Unique identifier for the relationship.
Constraints: Primary key, auto-generated
Index: Primary Key
```

**Column-2: Book ID**

```
Name: book_id
Data Type: UUID
Description: Foreign key referencing the "id" column in the Books table.
Constraints: NOT NULL, FK to Books(id)
Index: Foreign Key Index
```

**Column-3: Collection ID**

```
Name: collection_id
Data Type: UUID
Description: Foreign key referencing the "id" column in the Collections table.
Constraints: NOT NULL, FK to Collections(id)
Index: Foreign Key Index
```

## Relationships and Constraints

### Books -> Books_Collections:

- One book can belong to multiple collections
- *Relationship:* Books.id -> Books_Collections.book_id

### Collections -> Books_Collections:

- One collection can have multiple books
- *Relationship:* Collections.id -> Books_Collections.collection_id


