## Before Start

### Modules needed

Run the next commands

```
go get google.golang.org/grpc
```

## Command used to generate Protos

```
cd internal/adapters/grpc/proto
protoc --go_out=./go-src \
       --go-grpc_out=./go-src \
       --go_opt=paths=source_relative \
       --go-grpc_opt=paths=source_relative \
       order.proto
```

## Env Variables

```
export SERVER_PORT=8081
export ENV=dev

```

## Run Application

```
go run cmd/main.go
```

## Calling GRPC

```
// Simple RPC
grpcurl -d '{"item_name": "rice"}' -plaintext localhost:8081 Order/Create

// Server streaming RPC
grpcurl -d '{"orders_quantity": 60}' -plaintext localhost:8081 Order/ListOrders

// Client Streaming RPC
grpcurl -plaintext -d @ localhost:8081 Order/CreateStreamOrder <<EOF
{"item_name": "rice"}
{"item_name": "water"}
{"item_name": "juice"}
EOF

// Bidirectional Client/Server Streaming RPC
grpcurl -plaintext -d @ localhost:8081 Order/CreateBidirectionalStreamOrder <<EOF
{"item_name": "rice"}
{"item_name": "water"}
{"item_name": "juice"}
EOF

```

## Build docker image

```
GOOS=linux GOARCH=amd64 go build -o grpc-order-server ./cmd
```

## Useful Links

- GRPC Request types
  - https://grpc.io/docs/languages/go/basics/