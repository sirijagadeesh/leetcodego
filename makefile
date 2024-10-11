fmt:
	go fmt ./...

clean: fmt
	go clean -testcache

test: clean
	go test ./... -v