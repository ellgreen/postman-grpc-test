run: proto/*.pb.go
	go run .

protos: proto/*.pb.go

proto/*.pb.go: service.proto
	protoc --proto_path=. --go_out=. --go-grpc_out=. service.proto
