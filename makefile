build:
	go build -o project-markdown-builder ./cmd/cli/
test:
	go run ./cmd/cli/main.go