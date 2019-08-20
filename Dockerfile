FROM golang:1.12

# ENV TEST_ENV_LOAD="Hello World"

WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

COPY . ./

RUN go build -o bin/starter-api main.go

EXPOSE 8080

ENTRYPOINT [ "bin/starter-api" ]