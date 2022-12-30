compile:
	go build -o customers cmd/customers_cli.go

format:
	go fmt ./...

lint:
	go vet ./...
	staticcheck ./...

commit: compile format lint
