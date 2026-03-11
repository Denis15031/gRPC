# gRPC User & Chat API

Stub implementation of User and Chat gRPC services.

## Structure
- `proto/api.proto` - Protocol Buffer definitions
- `api/` - Generated code
- `internal/server/` - Service implementations
- `cmd/server/` - Main application

## Usage
```bash
./gen.sh          # Generate code
go build          # Build server
./server          # Run server