# Fake ARIA API

A mock API server that simulates VMware ARIA (formerly vRealize) API and vSphere responses for testing and development purposes. This project provides a lightweight server that mimics ARIA endpoints, allowing developers to test integrations with the [aria-vsphere-metrics-collector](https://github.com/Tata-Matata/aria-vsphere-metrics-collector) without requiring a real ARIA instance.

## Features

- **Deployments API** – Retrieve deployment information with randomized data
  - List all deployments: `GET /deployments`
  - Get specific deployment: `GET /deployments/{id}`
  
- **DataStores API** – Retrieve datastore information
  - List all datastores: `GET /datastores`
  - Get specific datastore: `GET /datastores/{id}`

- **Event Pushing** – Continuously publishes mock deployment events to an external exporter
  - Configurable event destination URL
  - Background event publisher that runs asynchronously

- **Structured Logging** – Comprehensive logging to file and console for debugging

## Architecture

The project is organized into the following packages:

- **`deploy`** – Handles deployment API endpoints and data management
- **`storage`** – Manages datastore API endpoints and data
- **`eventpusher`** – Publishes events to an external event exporter
- **`util`** – Utility functions for logging, JSON handling, and random data generation

## Prerequisites

- **Go** 1.22.6 or later

## Installation & Setup

Clone the repository:

```bash
git clone https://github.com/Tata-Matata/fake-aria-api.git
cd fake-aria-api
```

Install dependencies:

```bash
go mod download
```

Build the application:

```bash
go build -o fake-aria-api
```

## Running the Server

Run the compiled binary:

```bash
./fake-aria-api
```

Or run directly with Go:

```bash
go run main.go app.go constants.go
```

The server will start on `http://localhost:10000`

## Configuration

Server endpoints are configured in `constants.go`:

- **Fake API Server Port**: `:10000` (can be modified via `FAKE_API_SERVER_PORT`)
- **Event Exporter URL**: `http://localhost:11000/event` (can be modified via `EXPORTER_URL` and `EVENT_EXPORTER_ENDPOINT`)

## API Endpoints

### Deployments

**Get all deployments:**
```bash
curl http://localhost:10000/deployments
```

**Get a specific deployment:**
```bash
curl http://localhost:10000/deployments/{id}
```

### DataStores

**Get all datastores:**
```bash
curl http://localhost:10000/datastores
```

**Get a specific datastore:**
```bash
curl http://localhost:10000/datastores/{id}
```

## Usage with aria-vsphere-metrics-collector

This mock API is designed to be used as a test target for the [aria-vsphere-metrics-collector](https://github.com/Tata-Matata/aria-vsphere-metrics-collector) project:

1. Start the fake ARIA API:
   ```bash
   ./fake-aria-api
   ```

2. Configure the metrics collector to point to `http://localhost:10000` instead of a real ARIA instance

3. Run integration tests against this mock server

## Project Structure

```
fake-aria-api/
├── main.go              # Application entry point
├── app.go               # HTTP server setup and route handlers
├── constants.go         # Configuration constants
├── deploy/              # Deployment API implementation
│   ├── deploy.go
│   ├── deploy_test.go
│   └── constants.go
├── storage/             # DataStore API implementation
│   ├── storage.go
│   └── constants.go
├── eventpusher/         # Event publishing implementation
│   ├── event_pusher.go
│   ├── deploy-event-pusher.go
│   └── constants.go
├── util/                # Utility functions
│   ├── json.go
│   ├── log.go
│   ├── random.go
│   ├── maths.go
│   ├── util_test.go
│   └── constants.go
├── data/                # Mock data
│   └── deployments.json
└── go.mod              # Go module definition
```

## Testing

Run unit tests:

```bash
go test ./...
```

Run tests with verbose output:

```bash
go test -v ./...
```


## Future Enhancements

See [TODO.md](./TODO.md) for planned improvements, including:

- API endpoint tests
- Health check endpoint
- Configuration file support
- CI/CD pipeline integration

