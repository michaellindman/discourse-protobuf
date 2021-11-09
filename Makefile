update-debs:
	go get -u ./...

proto:
	protoc \
	--proto_path=/usr/local/include/google/protobuf \
	--proto_path=proto/ \
	--go_out=$(GOPATH)/src proto/*.proto

.PHONY: proto
