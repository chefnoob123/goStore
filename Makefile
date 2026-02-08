build:
	@go build -o bin/gS

run: build
	@./bin/gS

test:
	@go test ./... -v
