proto:
	protoc --go_out=. --go_opt=paths=source_relative \
			--go-grpc_out=. --go-grpc_opt=paths=source_relative \
			internal/adapters/grpc/proto/*.proto

.PHONY: proto