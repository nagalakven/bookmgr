
# REST API Documentation

## Overview

The Book Management System application provides a set of RESTful APIs for interaction over HTTP. All API calls and responses are handled over standard HTTP methods (GET, POST, PUT, DELETE) and JSON data format.

TODO: This this application does not support authentication at present, so all endpoints are accessible without using API keys or other authentication mechanisms.

## Base URL

The base URL for accessing the API is:

```
http://localhost:8080/api
```

All requests should be made relative to this base URL.

TODO: Localhost is used for development purpose. Actual hostname and domain needs to be updated.

## API Versioning

This API follows simple versioning where the major version is specified in the URL. The current version is:

```
/api/v1/
```

### Supported API Versions

You can retrieve the available versions of the API using the following endpoint:

```
GET /api/versions
```

This will return a list of supported versions.

## Return values

These are the standard return types:

* Standard return value
* Error

## Standard return value

For a standard synchronous operation, the following JSON object is returned:

```js
{
    "data": {}          //free form response based on the operation
}
```

HTTP code must be 200.

## Error Handling

If an error occurs during an API request, the response will contain an error response with the details of the issue.

### Example Error Response

```json
{
    "error": "<error message>"
}
```

### Common Error Codes

| Code | Meaning                    |
|------|----------------------------|
| 400  | Bad Request                |
| 404  | Resource Not Found         |
| 500  | Internal Server Error      |

## Status Codes

- **200**: OK – The request was successful.
- **201**: Created - The resource is successfully created.
- **400**: Bad Request – Invalid request format or parameters.
- **404**: Not Found – The resource was not found.
- **500**: Internal Server Error – An unexpected error occurred on the server.

## Pagination

If the result set of a query is large, pagination can be used to retrieve the data in smaller chunks. 

Use the `limit` query parameters to paginate through results.

### Example

```http
GET /api/v1/books?limit=10
```

TODO: Implementation to be done

## Asynchronous Operations

TODO: Place holder for future implementation of asynchronous operations

## API structure 

The YAML version of this API specification can be found in [`rest-api.yaml`](./rest-api.yaml)
