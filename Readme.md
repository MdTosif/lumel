Lumel API Documentation
=======================

Introduction
------------

The Lumel API provides endpoints to manage orders and retrieve insights into product sales based on various criteria. Built using the Gin web framework and GORM for database interactions, it offers functionalities for importing orders and querying top-selling products.

Project Setup
-------------

To set up the project, change the database credentials in `./config/database.go`. (Note: There hasn't been time to separate this configuration into a dedicated file.)

postman collection docs link: https://documenter.getpostman.com/view/12703692/2sAXxMftZz
Endpoints
---------

### 1\. Import Orders from CSV

**Endpoint:** `POST /orders/import`  
**Description:** Imports orders from a CSV file. If a customer, order, or product does not exist, it will be created. If an order already exists with the same ID, it will be skipped.

**Request:** A form file containing the CSV data.

**Response:**

*   `200 OK:` If the file is processed successfully.
*   `400 Bad Request:` If no file is provided.
*   `500 Internal Server Error:` If there is an error processing the file.

### 2\. Top N Products Overall

**Endpoint:** `POST /orders/top-overall`  
**Description:** Retrieves the top N products based on quantity sold within a specified date range.

**Request Body:**

    {
        "limit": 5,
        "start_date": "2023-12-15T03:30:00+05:30",
        "end_date": "2023-12-15T09:30:00+05:30"
    }

**Example cURL Request:**

    curl --location 'localhost:8000/orders/top-overall' \
    --data '{
        "limit": 5,
        "start_date": "2023-12-15T03:30:00+05:30",
        "end_date": "2023-12-15T09:30:00+05:30"
    }'

**Response:**

*   `200 OK:` Returns a list of top N products.

**Example Response:**

    [
      {
        "product_id": "P123",
        "quantity_sold": 2
      }
    ]

### 3\. Top N Products by Region

**Endpoint:** `POST /orders/top-by-region/:region`  
**Description:** Retrieves the top N products sold in a specified region within a given date range.

**Request Body:**

    {
        "limit": 5,
        "start_date": "2023-12-15T03:30:00+05:30",
        "end_date": "2023-12-15T09:30:00+05:30"
    }

**Example cURL Request:**

    curl --location 'localhost:8000/orders/top-by-region/Anytown' \
    --data '{
        "limit": 5,
        "start_date": "2023-12-15T03:30:00+05:30",
        "end_date": "2023-12-15T09:30:00+05:30"
    }'

**Response:**

*   `200 OK:` Returns a list of top N products in the specified region.

**Example Response:**

    [
      {
        "product_id": "P123",
        "quantity_sold": 2
      }
    ]

### 4\. Top N Products by Category

**Endpoint:** `POST /orders/top-by-category/:category`  
**Description:** Retrieves the top N products sold in a specified category within a given date range.

**Request Body:**

    {
        "limit": 5,
        "start_date": "2023-12-15T03:30:00+05:30",
        "end_date": "2023-12-15T09:30:00+05:30"
    }

**Example cURL Request:**

    curl --location 'localhost:8000/orders/top-by-category/Shoes' \
    --data '{
        "limit": 5,
        "start_date": "2023-12-15T03:30:00+05:30",
        "end_date": "2023-12-15T09:30:00+05:30"
    }'

**Response:**

*   `200 OK:` Returns a list of top N products in the specified category.

**Example Response:**

    [
      {
        "product_id": "P123",
        "quantity_sold": 2
      }
    ]

Notes
-----

*   Ensure your database is configured correctly in the config package before running the application.
*   The API runs on `localhost:8000`.