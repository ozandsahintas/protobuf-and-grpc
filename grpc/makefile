PROTO_DIR = proto

ifeq ($(OS), Windows_NT)
	PACKAGE = $(shell (Get-Content go.mod -head 1).Split(" ")[1])
else
	PACKAGE = $(shell head -1 go.mod | awk '{print $$2}')
endif

.PHONY: initial unary server-streaming client-streaming bi-directional
project := initial unary server-streaming client-streaming bi-directional

$(project):
	protoc --go_opt=module=${PACKAGE}/$@ --go_out=$@/${PROTO_DIR} --go-grpc_opt=module=${PACKAGE}/$@ --go-grpc_out=$@/${PROTO_DIR} $@/${PROTO_DIR}/*.proto