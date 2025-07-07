# gRPC Network Prober
A simple network monitoring tool that checks website response times using gRPC.
---
## What This Does
- Server: Makes HTTP requests to websites and measures response times
- Client: Asks the server to check specific websites and reports back the results

## Getting Started
### Install Dependencies

```bash
go mod tidy
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
### Generate Code from Proto Files
```bash
protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    prober/prober.proto
``` 
## How to Use
### Start the Server
```bash
go run prober-server/main.go
```
### Run the Client
```bash
go run prober_client/main.go
```

## Project Structure
grpc-prober/
├── prober/
│   ├── prober.proto          # API definition
│   ├── prober.pb.go          # Generated Go code
│   └── prober_grpc.pb.go     # Generated Go code
├── prober_server/
│   └── main.go               # Server implementation
├── prober_client/
│   └── main.go               # Client implementation
└── go.mod                    # Go dependencies