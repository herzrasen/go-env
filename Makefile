all: clean test

clean:
	rm -f coverage.out

test:
	go test -coverprofile=coverage.out -v ./...
	go tool cover -func coverage.out
