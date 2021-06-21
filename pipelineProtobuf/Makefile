SHELL := /bin/bash

protos: ./data.proto
	protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative data.proto

.PHONY: clean-protos
clean-protos:
	rm -f data.pb.go