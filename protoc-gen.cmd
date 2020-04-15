protoc --proto_path=api/proto/v1 --proto_path=third_party --go_out=plugins=grpc:api/proto/v1 test.proto
