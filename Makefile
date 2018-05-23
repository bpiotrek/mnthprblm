run: main.go
	go run *.go

format:
	gofmt -w -s -d *.go

.PHONY: run format
