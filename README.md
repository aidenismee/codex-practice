# Echo Starter

A minimal Go API starter built with the [Echo](https://echo.labstack.com/) framework.

## Requirements

- Go 1.21+

## Getting Started

Install dependencies (only needed the first time or after updates):

```bash
go mod download
```

Run the server:

```bash
go run ./cmd/server
```

The server listens on `http://localhost:8080`.

## API Endpoints

- `GET /` returns a friendly message.
- `GET /health` returns a JSON status payload.

### Example Requests

```bash
curl http://localhost:8080/
curl http://localhost:8080/health
```

## Project Structure

- `cmd/server` - application entrypoint.
- `internal/routes` - HTTP route registration.
- `internal/handlers` - request handlers.

## Development Notes

- Logging and recovery middleware are enabled by default.
- Update routes in `internal/routes/routes.go` to add new endpoints.
