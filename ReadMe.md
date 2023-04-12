# Message Queuing Application (RabbitMQ + MongoDB)

This project is a Message Queuing application that demonstrates the use of RabbitMQ and MongoDB for processing and storing product data. The application is divided into a producer, a consumer, and a utility for adding or updating user data.

## Functionality

The application's main goal is to process product data in the form of images by downloading, compressing, and storing images associated with the products. The producer exposes an API for adding new products and publishes messages to a **RabbitMQ queue**. The consumer listens for messages in the queue, processes the product data as required, and stores the data in a **MongoDB database**.

**RabbitMQ** has been chosen over Kafka for this project due to its simplicity, and ease of use.

## Project Structure

The project is organized into the following directories:

- **`cmd`:** Contains the main entry points for the various executables of the application, including the `addUser` utility, the `consumer` application, and the `producer` application.     
- **`internal`:** Contains the internal packages used by the executables, such as `config`, `db`, `models`, `rabbitmq`, and the `app` package for the producer and consumer applications.     

For detailed information on each directory, refer to their respective `README.md` files.

## How to Get Started?

1. Ensure that you have [Go](https://golang.org/doc/install) installed on your system.
2. Also make sure to have MongoDB and RabbitMQ running on your machine.(You can make use of docker containers to run locally on your machine)
3. Clone the repository: `git clone https://github.com/SinisterSup/messageQueuing.git`
4. Navigate to the project root directory: `cd messageQueuing`
5. Install the dependencies: `go mod download`
6. Create a `config.yaml` file in the root directory with the necessary configuration for MongoDB and RabbitMQ. You can use the `config.example.yaml` file as a reference.
7. Start the `producer` and `consumer` applications by following the instructions in the `cmd` directory's `README.md` file.
8. Optionally, you can run the `addUser` utility to add or update user data from a JSON file by following the instructions in the `cmd` directory's `README.md` file.

---
