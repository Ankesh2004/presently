# There are 2 images: "build" is used for the CI/CD pipeline and "run" is what actually gets deployed.

FROM golang:1.12 AS build

ENV PORT=8080

# Install golangci-lint. This is only needed in the build image
# https://github.com/golangci/golangci-lint#ci-installation
RUN wget -O - -q https://install.goreleaser.com/github.com/golangci/golangci-lint.sh | sh -s latest

# Install wait-for-it which is used to wait for cassandra to finish starting up before
# executing the tests
RUN apt-get update && apt-get install -y "wait-for-it"

WORKDIR /app

# Pre-run the go deps in order to possible get caching on subsequent container builds
COPY ./go.* ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 make build

FROM alpine:3.7 AS run
COPY --from=build /app/bin/starter-api /bin/starter-api
ENV PORT=8080
EXPOSE 8080

ENTRYPOINT [ "/bin/starter-api" ]
