generate:
	echo "Generating files..."
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

run:
	go run main.go
