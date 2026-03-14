.PHONY: test test-v test-exercises test-assignments cover clean

test:
	go test ./...

test-v:
	go test -v ./...

test-exercises:
	go test ./exercises/...

test-homework:
	go test ./homework/...

cover:
	go test -cover ./...

clean:
	go clean
