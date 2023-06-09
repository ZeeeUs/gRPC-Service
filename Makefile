generate_server:
	@echo "Generating Go files from Protobuf definitions..."
	@protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	internal/domain/proto/auto_market.proto

generate_gateway:
	@echo "Generating Gateway files from Protobuf definitions..."
	@protoc -I . \
    --grpc-gateway_out=. \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    internal/domain/proto/auto_market.proto