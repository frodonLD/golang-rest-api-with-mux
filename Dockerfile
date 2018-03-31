
# Build stage
FROM golang:1.10-alpine3.7 as builder

RUN apk --no-cache add -U make git musl-dev gcc

WORKDIR $GOPATH/src/github.com/frodonLD/GoLangRESTAPIWithMux
ADD . .

# Install dependencies
RUN go get -u github.com/golang/dep/cmd/dep
RUN dep ensure

# Run the tests and benchmarks
RUN go test ./...

# Install app
RUN go install


# Final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /app
COPY --from=builder /go/bin/GoLangRESTAPIWithMux .

ENTRYPOINT ["/app/GoLangRESTAPIWithMux"]