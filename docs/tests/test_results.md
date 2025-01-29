# Test Results
This document outlines the curl requests and their corresponding responses for some unit test cases.

## Step 1: POST request to add a new book with only mandatory fields

```
curl -X POST http://localhost:8080/api/books -H "Content-Type: application/json" -d "{\"Title\": \"Go Programming\", \"Author\": \"Author-1\", \"Published_Date\": \"2025-01-01\"}" 
```

**Response:**

```
{"data":{"book":{"id":"1250c637-4987-442d-bb85-45b0c2837b82","title":"Go Programming","author":"Author-1","published_date":"2025-01-01","created_at":"2025-01-28T22:40:09.8671765-08:00","updated_at":"2025-01-28T22:40:09.8671765-08:00"},"id":"1250c637-4987-442d-bb85-45b0c2837b82","message":"Book added successfully!"}}
```

## Step 2: POST request to add a new book with all fields

```
curl -X POST http://localhost:8080/api/books -H "Content-Type: application/json" -d "{\"Title\": \"Bookmgr\", \"Author\": \"Naga\", \"Published_Date\": \"2022-01-01\", \"published_by\": \"Naga\", \"co_authors\": \"co-1,co-2,co-3\", \"genre\": \"thriller\", \"description\": \"a thriller book\", \"edition\": \"21\", \"recommended_age\": \"10+\", \"is
bn\": \"4314-434-243-432\"}"
```

**Response:**

```
{"data":{"book":{"id":"14b14765-166c-470f-be57-a06586cc30d1","title":"Bookmgr","author":"Naga","published_date":"2022-01-01","published_by":"Naga","co_authors":"co-1,co-2,co-3","genre":"thriller","description":"a thriller book","edition":"21","recommended_age":"10+","isbn":"4314-434-243-432","created_at":"2025-01-28T22:43:46.5754829-08:00","updated_at":"2025-01-28T22:43:46.5754829-08:00"},"id":"14b14765-166c-470f-be57-a06586cc30d1","message":"Book added successfully!"}}
```

## Step 3: GET request to list all books

```
curl -X GET http://localhost:8080/api/books -H "Content-Type: application/json"
```

**Response:**

```
{"data":{"books":[{"id":"14b14765-166c-470f-be57-a06586cc30d1","title":"Bookmgr","author":"Naga","published_date":"2022-01-01","published_by":"Naga","co_authors":"co-1,co-2,co-3","genre":"thriller","description":"a thriller book","edition":"21","recommended_age":"10+","isbn":"4314-434-243-432","created_at":"2025-01-28T22:43:46.5754829-08:00","updated_at":"2025-01-28T22:43:46.5754829-08:00"},{"id":"1250c637-4987-442d-bb85-45b0c2837b82","title":"Go Programming","author":"Author-1","published_date":"2025-01-01","created_at":"2025-01-28T22:40:09.8671765-08:00","updated_at":"2025-01-28T22:40:09.8671765-08:00"}],"message":"List of Books"}}
```

## Step 4: PUT request to update an existing book (non-mandatory fields)

```
curl -X PUT http://localhost:8080/api/books/14b14765-166c-470f-be57-a06586cc30d1 -H "Content-Type: application/json" -d "{\"Title\": \"Bookmgr\", \"Author\": \"Naga\", \"Published_Date\": \"2022-01-01\", \"published_by\": \"Naga\", \"co_authors\": \"co-1,co-2,co-3,co-4\", \"genre\": \"thriller\", \"description\": \"a thriller book\", \"recommended_age\": \"12+\"}" 
```

**Response:**

```
{"data":{"id":"14b14765-166c-470f-be57-a06586cc30d1","message":"Book updated successfully!","updated book":{"id":"14b14765-166c-470f-be57-a06586cc30d1","title":"Bookmgr","author":"Naga","published_date":"2022-01-01","published_by":"Naga","co_authors":"co-1,co-2,co-3,co-4","genre":"thriller","description":"a thriller book","recommended_age":"12+","created_at":"2025-01-28T22:43:46.5754829-08:00","updated_at":"2025-01-28T22:46:39.5405554-08:00"}}}
```

## Step 5: DELETE request to delete an existing book

```
curl -X DELETE http://localhost:8080/api/books/1250c637-4987-442d-bb85-45b0c2837b82 -H "Content-Type: application/json" 
```

**Response:**

```
{"data":{"message":"Book deleted successfully!"}}
```

## Step 6: GET request to verify with list all books

```
curl -X GET http://localhost:8080/api/books -H "Content-Type: application/json"
```

**Response:**

```
{"data":{"books":[{"id":"14b14765-166c-470f-be57-a06586cc30d1","title":"Bookmgr","author":"Naga","published_date":"2022-01-01","published_by":"Naga","co_authors":"co-1,co-2,co-3,co-4","genre":"thriller","description":"a thriller book","recommended_age":"12+","created_at":"2025-01-28T22:43:46.5754829-08:00","updated_at":"2025-01-28T22:46:39.5405554-08:00"}],"message":"List of Books"}}
```

## Step 7: POST request to add a new book with a mandatory field missing

```
curl -X POST http://localhost:8080/api/books -H "Content-Type: application/json" -d "{\"Title\": \"Book-1\", \"Author\": \"Author-1\", \"genre\": \"fiction\"}"
```

**Response:**

```
{"error":"validation failed: title, author and published date are required"}
```

## Step 8: PUT request to update mandatory fields of an existing book

```
curl -X PUT http://localhost:8080/api/books/14b14765-166c-470f-be57-a06586cc30d1 -H "Content-Type: application/json" -d "{\"Title\": \"Book-2\", \"Author\": \"Author-2\", \"published_date\": \"1990-04-10\"}" 
```

**Response:**

```
{"data":{"id":"14b14765-166c-470f-be57-a06586cc30d1","message":"Book updated successfully!","updated book":{"id":"14b14765-166c-470f-be57-a06586cc30d1","title":"Book-2","author":"Author-2","published_date":"1990-04-10","created_at":"2025-01-28T22:43:46.5754829-08:00","updated_at":"2025-01-28T22:52:40.3224857-08:00"}}}
```

## Step 9: GET request to verify with list all books

```
curl -X GET http://localhost:8080/api/books -H "Content-Type: application/json"
```

**Response:**

```
{"data":{"books":[{"id":"14b14765-166c-470f-be57-a06586cc30d1","title":"Book-2","author":"Author-2","published_date":"1990-04-10","created_at":"2025-01-28T22:43:46.5754829-08:00","updated_at":"2025-01-28T22:52:40.3224857-08:00"}],"message":"List of Books"}}
```

## Step 10: DELETE request to delete a non existing book

```
curl -X DELETE http://localhost:8080/api/books/91d43740-ed13-4d07-9514-bedf61857511 -H "Content-Type: application/json" 
```

**Response:**

```
{"error":"Book [91d43740-ed13-4d07-9514-bedf61857511] not found."}
```

## Step 11: POST request with invalid payload

```
curl -X POST http://localhost:8080/api/books -H "Content-Type: application/json" -d "{\"Titfedfawf 
```

**Response:**

```
{"error":"invalid request payload"}
```
