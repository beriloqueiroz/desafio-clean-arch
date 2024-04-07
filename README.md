# Challenge Execution Guide

This guide details the steps necessary to prepare and execute the components of the clean architecture challenge. Follow the instructions below to test the REST and GraphQL APIs, as well as the gRPC API.

## Environment Setup

Before starting the tests, it's necessary to prepare the environment with the following steps:

1. **Start Services with Docker Compose**

    Run the command below to start all the necessary services using Docker Compose. This command should be run at the project root:

    ```bash
    docker compose up -d
    ```

## Testing the APIs

### REST API

To test the REST API, use the HTTP files available in the `/api` directory. Follow the instructions below:

- **Create Order**: Use the file `/api/create_order.http` to create a new order.
- **List Orders**: Use the file `/api/list_order.http` to list all existing orders.

### GraphQL API

To test the GraphQL API, follow the steps below:

1. Copy the content of the file `/api/mutations.graph`.
2. Access the GraphQL graphical interface at [http://localhost:8080/manager/html/](http://localhost:8080/manager/html/).
3. Paste the copied content into the graphical interface and execute the query or mutation.

### gRPC API

To test the gRPC API, use the Evans tool. Run the command below at the project root:

    ```bash
    evans --host 127.0.0.1 ./internal/infra/grpc/protofiles/order.proto
    ```

### Checking Records in RabbitMQ

To check the records in RabbitMQ, access the following address in your browser:

[http://localhost:15672/#/queues](http://localhost:15672/#/queues)

This link will lead to the RabbitMQ interface, where you can check the queues and messages.
