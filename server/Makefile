generate:
	protoc -I. \
  --go_out=internal/generated/go \
  --go-grpc_out=internal/generated/grpc \
  --grpc-gateway_out=logtostderr=true:. \
  --openapiv2_out ./api \
  server/api/go_load.proto
