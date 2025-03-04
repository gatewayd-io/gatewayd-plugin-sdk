lint:
	@buf lint

gen:
	@buf generate

update-all:
	go get -u ./...
	go mod tidy

install-deps:
	@go install github.com/bufbuild/buf/cmd/buf@latest
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install	google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
	@go install	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest
	@go install	github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest

# https://groups.google.com/g/golang-nuts/c/FrWNhWsLDVY/m/CVd_iRedBwAJ
update-direct-deps:
	@go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all | xargs -n1 go get
	@go mod tidy
