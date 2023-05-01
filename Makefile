generate_server:
	protoc \
	--go_out=. \
	--go_opt=paths=source_relative \
	--go-grpc_out=. \
	--go-grpc_opt=paths=source_relative \
	internal/domain/proto/social_network.proto

generate_gateway:
	protoc -I . \
    --grpc-gateway_out=. \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    internal/domain/proto/social_network.proto
