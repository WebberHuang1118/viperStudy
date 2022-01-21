SHELL := /bin/bash

protos: data/data.proto
	protoc --go_out=plugins=grpc:. --go_opt=paths=source_relative data/data.proto	

.PHONY: clean-protos
clean-protos:
	rm -f data/data.pb.go