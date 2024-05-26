# Exchange Rate Receiver

## Overview

Exchange Rate Receiver is a GRPC service that retrieves the USDT exchange rate from the Garantex exchange. It provides ask and bid prices along with a timestamp, and stores this information in a PostgreSQL database upon each request.

## Features

- GRPC method `GetRates` for retrieving USDT exchange rates from Garantex.
- Stores the exchange rates with a timestamp in PostgreSQL.
- Healthcheck method to check the service's health.
- Graceful shutdown.

## Requirements

- Go 1.22 or higher
- Docker and Docker Compose
- PostgreSQL

## Configuration

Configuration parameters can be managed using environment variables or command-line flags. The following parameters are used:

- `DATABASE_URL`: PostgreSQL connection string (`postgres://user:password@localhost:5432/exchangerates?sslmode=disable`)
- `GRPC_PORT`: Port for the GRPC server (default: `50051`)

Copy .env.example to .env in your local environment

## Setup Instructions

### 1. Clone the repository

```sh
git clone https://github.com/yourusername/exchange-rate-receiver.git
cd exchange-rate-receiver
```

### 2. Build the application

```bash
make build
```

### 3. Build the Docker images and run the services using Docker Compose
```bash
docker-compose build
docker-compose up
```

### 4. Run the application
```bash
docker-compose run --rm app ./app
```
### 5. Run the tests
```bash
make test
```

## Protobuf
### Install
(MacOS)
To install `protoc` run
```bash
brew install protobuf
```
Then download Go protobuf plugins
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
### Generation
Navigate to the root directory of your project and run the following commands to generate the Go code from the protobuf definitions:

```bash
protoc --go_out=.. --go-grpc_out=.. proto/exchange_rate.proto
```

### Running Locally

1. Install dependencies:
   ```sh
   go mod download
   ```

2. Build the application:
   ```sh
   cd cmd/server
   go build -o ../../exchange_rate-receiver
   cd ../..
   ```

3. Run the application:
   ```sh
   ./exchange_rate-receiver
   ```

