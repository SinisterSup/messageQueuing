# App

This folder contains the core application logic for the message queuing system. It is divided into two sub-packages: `api` and `consumer`.

## API

The `api` package handles the server side of the application. It processes the **REST**ful API for clients to interact with the message queuing system. The main functionality provided by this package is to create new products and send them to a message queue (**RabbitMQ**) for further processing.

The `api` package contains the following files:

- `handlers.go`: This file contains the main API handler functions, such as `CreateProduct`, which is responsible for creating a new product, inserting it into the database, and sending the product ID to the RabbitMQ queue.

- `routes.go`: This file defines the routes for the API. Currently, it only has one route: a **POST** request to `/api/products`, which triggers the `CreateProduct` function from `handlers.go`.

## Consumer

The `consumer` package is responsible for processing messages from the RabbitMQ queue. Specifically, it handles the task of downloading, compressing, and storing product images.

The `consumer` package contains the following files:

- `consumer.go`: This file contains the main logic for the consumer. It includes functions like `DownloadCompressAndStoreImages`, which is responsible for downloading and compressing the images of a product, and updating the product entry in the database with the compressed images.

## Usage

To use the API, make a POST request to `/api/products` on `http://localhost:8080` with the following JSON body:

```json
{
  "user_id": "example_user_id",
  "product_name": "Example Product",
  "product_description": "This is an example product.",
  "product_images": [
    "https://example.com/image1.jpg",
    "https://example.com/image2.jpg"
  ],
}
