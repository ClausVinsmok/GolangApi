.PHONY:build

build:
	GOOS=windows GOARCH=amd64 go build -o ./apiserver ./cmd/apiserver/main.go

.PHONY:test
test:
	go test -v -race -timeout 30s ./...

.DEFAULT_GOAL:= build