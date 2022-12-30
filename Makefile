format:
	go fmt ./...

dependencies:
	go mod tidy

compile:
	go build ./...

test_all:
	go test -count=1 -cover ./...

commit: format dependencies compile
