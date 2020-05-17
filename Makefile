.PHONY: proto
proto:
	#everyline new after this should be on tab
	protoc \
		-I pkg/proto/ \
        -I ${GOPATH}/src \
        --go_out=plugins=grpc:pkg/proto \
        pkg/proto/api.proto;

.PHONY: build-server
build-server:
	go build -i -v -o bin/server github.com/atyagi9006/hello-grpc/server

.PHONY: build-client
build-client:
	go build -i -v -o bin/client github.com/atyagi9006/hello-grpc/client