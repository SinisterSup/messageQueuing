# Internal Package

The `internal/` package contains the core components of the message queuing application, organized into several sub-packages. Each sub-package handles a specific functionality and is designed to be modular and easy to understand.
    
internal/    
│    
├── app/    
│ ├── api/  &emsp;  # API-related handlers and routes    
│ └── consumer/  &emsp;  # Consumer logic for processing messages from the queue    
│     
├── config/  &emsp;  # Configuration settings for the application    
│    
├── db/  &emsp;  # Database connection and related functionality for MongoDB    
│    
├── models/  &emsp;  # DataBase schema models for the application    
│     
└── rabbitmq/  &emsp;  # RabbitMQ connection and related functionality    

         
    
## Sub-packages Overview

### app/api

This package contains the API-related handlers and routes. It is responsible for handling incoming HTTP requests, processing them, and returning appropriate responses.

### app/consumer

This package contains the logic for consuming messages from the RabbitMQ queue. It listens for incoming messages and processes them accordingly.

### config

This package contains the configuration settings for the application. It defines a `Config` struct, which holds the MongoDB and RabbitMQ configuration details, and a `LoadConfig()` function that initializes the configuration.

### db

This package contains the functions for connecting to a MongoDB database and retrieving a specific collection. It exports two functions: `ConnectMongoDB()` for connecting to a MongoDB instance, and `GetCollection()` for retrieving a specific collection from the connected MongoDB database.

### models

This package contains the data models for the application. It defines the `User` and `Product` structs, which represent the structure of the data stored in the database.

### rabbitmq

This package contains the functions for connecting to a RabbitMQ instance and sending messages to a queue. It exports two functions: `ConnectRabbitMQ()` for connecting to a RabbitMQ instance, and `SendProductID()` for sending a product ID to the RabbitMQ queue.

---

To use the internal packages, simply import the required sub-package in your application and make use of the exported functions and data structures.