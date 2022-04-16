OUT=./pb/calc
NAME=calc.proto

.PHONY: protoc
protoc:
	protoc --proto_path ./proto --go_out=plugins=grpc:${OUT} ${NAME}
