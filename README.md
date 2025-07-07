# gRPC Network Prober
A simple network monitoring tool that checks website response times using gRPC.
---
## What This Does
Server: Makes HTTP requests to websites and measures response times
Client: Asks the server to check specific websites and reports back the results

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
