# Makefile

# Make sure that the go lang bin is included in you PATH then this is called.
protoc:
	cd proto && protoc \
	--go_out=../proto --go_opt=paths=source_relative \
	--go-grpc_out=../proto --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=../proto --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
	./**/*.proto

protoc-plus:
	cd proto && protoc \
	--include_imports \
    --include_source_info \
	--go_out=../proto --go_opt=paths=source_relative \
	--go-grpc_out=../proto --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=../proto --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    --descriptor_set_out=api_descriptor.pb \
	./**/*.proto