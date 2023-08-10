# Simple Go CRUD 

## TechStack
- GIN
- GORM (SQLite)

## Table of Contents

- [Installation](#installation)
- [API Endpoints](#api-endpoints)
  - [Readiness Check](#readiness-check)
  - [Create Customer](#create-customer)
  - [Update Customer](#update-customer)
  - [Get Customer by ID](#get-customer-by-id)
  - [Delete Customer](#delete-customer)
- [Running Unit Tests](#running-unit-tests)

## Installation

To get started with the **Project Name** project, follow these steps:

1. **Clone the Repository:**
   ```
   git clone https://github.com/BoomNooB/go-practice-sql.git
   cd go-practice-sql
   ```

3. **Install Dependencies:**\
   Ensure you have Go installed. Then, run the following command to install the project's dependencies:\
   ```go mod tidy```

4. **Create .env File:**\
   Create a `.env` file in the root directory and set the \`API_PORT\` variable to specify the port on which the API will run:
   Assume you'll use port 22345 so it will be
   ```API_PORT=22345```

6. **Initialize the Database:**\
   Run the following command to initialize the SQLite database and put 20 example of customer info into database:
   `go run createInitDataInDB.go`

7. **Run the Application:**\
   Start the server by executing:
   `go run cmd/main/main.go`\
   or you can run with `air` (i've already provide the air config file)
   The server will start at `http://localhost:22345`

## API Endpoints

### Readiness Check

Check if the API is running.

- **URL:** `/`
- **Method:** `GET`
- **Response:**
  - Status Code: 200 OK
  - Body:
```
    {
      "data": "API is running"
    }
```
### Create Customer

Create a new customer.

- **URL:** `/customers`
- **Method:** `POST`
- **Request Body:**
```
  {
    "name": "New Customer",
    "age": 25
  }
```

- **Response:**
  - Status Code: 201 Created
  - Body:
```
    {
      "id": 2,
      "name": "New Customer",
      "age": 25
    }
```

### Update Customer

Update customer information by ID.

- **URL:** `/customers/:id`
- **Method:** `PUT`
- **URL Parameters:**
  - `id` (uint, required) - The ID of the customer to update.
- **Request Body:**
```
  {
    "name": "Updated Customer",
    "age": 30
  }
```

- **Response:**
  - Status Code: 200 OK
  - Body:
```
    {
      "id": 1,
      "name": "Updated Customer",
      "age": 30
    }
```

### Get Customer by ID

Retrieve customer information by ID.

- **URL:** `/customers/:id`
- **Method:** `GET`
- **URL Parameters:**
  - `id` (uint, required) - The ID of the customer to retrieve.
- **Response:**
  - Status Code: 200 OK
  - Body:
```
    {
      "id": 1,
      "name": "John Doe",
      "age": 30
    }
```

### Delete Customer

Delete a customer by ID.

- **URL:** `/customers/:id`
- **Method:** `DELETE`
- **URL Parameters:**
  - `id` (uint, required) - The ID of the customer to delete.
- **Response:**
  - Status Code: 200 OK

## Running Unit Tests

Before running the unit tests, follow these steps:

1. Run the following command to generate test data:
   `go run createInitDataInDB.go`
   

3. Copy the generated `customer.db` file to the `./api` folder.

4. In the test file `api/customer_test.go`, ensure that the variable `record_id` contains an existing ID from the database.

5. Copy the `.env` file to the `./api` folder.

6. Run the tests using the following command:
   `go test ./...`

The test coverage of this suite is **83.9%.**

---
