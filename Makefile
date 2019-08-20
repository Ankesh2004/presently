build-image: build
	docker build -t starter-api .

build: clean
	env go build -o bin/starter-api  main.go

clean:
	rm -rf ./bin

test:
	go test -v ./...

dev:
	# Using gin for auto rebuilding binary when source code changes
	# https://github.com/codegangsta/gin
	# Install with `go get github.com/codegangsta/gin`
	# Make sure that GOPATH is setup. On OSX see:
	# https://ahmadawais.com/install-go-lang-on-macos-with-homebrew/
	gin --port 9998 --appPort 9999 run ./main.go 9999

