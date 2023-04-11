# CMD Directory

The `cmd` directory contains the main entry points for the various executables of the message queuing application.

## Structure

The directory is organized as follows:

- `addUser`: A utility to add or update users from a JSON file (`users.json`).
- `consumer`: The main consumer application responsible for consuming messages from the **RabbitMQ** queue and processing them.
- `producer`: The main producer application which posts an API for adding new products and publishing messages to the RabbitMQ queue.
     
     
## consumer
      
The consumer application connects to the **RabbitMQ** server and consumes messages from the `product_queue`. Upon receiving a message, it processes the message by downloading, compressing, and storing images for the corresponding product.         
To run the consumer application, navigate to the `consumer` directory and execute: `go run main.go`.
     
     
## producer
     
The producer application establishes a connection to POST an API for adding new products and publishes messages to the RabbitMQ queue. The API listens on port `8080` (generally referred to `http://localhost:8080`).        
To run the producer application, navigate to the `producer` directory and execute: `go run main.go`.
     
    
### addUser

This util directory reads user data from a JSON file (`users.json`) and either adds new users or updates existing ones in the **MongoDB** database.       
To run this utility, navigate to the `adduser` directory and execute: `go run main.go`. 
     
The `users.json` file should contain an array of user objects with the following properties:

- `user_id`: The primary key i.e, unique identifier of the user.
- `name`: The name of the user.
- `mobile`: The mobile number of the user.
- `latitude`: The latitude of the user's location.
- `longitude`: The longitude of the user's location.