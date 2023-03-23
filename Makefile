lint:
	@buf lint

gen:
	@buf generate

update-all:
	go get -u ./...
	go mod tidy
