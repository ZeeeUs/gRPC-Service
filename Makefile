generate_server:
	protoc \
	--go_out=pkg/social_network \
	--go_opt=paths=source_relative \
	--go-grpc_out=pkg/social_network \
	--go-grpc_opt=paths=source_relative \
	proto/social_network.proto