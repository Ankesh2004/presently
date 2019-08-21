.PHONY: build run dev test

build-image: build
	docker build -t starter-api .

clean:
	rm -rf ./bin

test:
	go test -v ./...

build: clean
	go build -o ./bin/starter-api ./main.go

run: build
	./bin/starter-api 9999

# Using reflex to watch for changes to .go file
# and re-run `make run`
# https://github.com/cespare/reflex/issues/50#issuecomment-388099690
# Install with `go get github.com/cespare/reflex`
# Make sure that GOPATH is setup. On OSX see:
# https://ahmadawais.com/install-go-lang-on-macos-with-homebrew/
dev:
	reflex --start-service -r '\.go$$' make run
