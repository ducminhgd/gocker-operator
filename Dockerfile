FROM golang:1.15.6-alpine3.12 AS builder
ENV GO111MODULE=on \
    CGO_ENABLED=1  \
    GOARCH="amd64" \
    GOOS=linux
WORKDIR /app
COPY . .
RUN apk add --no-cache git && go mod download && go build -ldflags="-s -w" -tags musl --ldflags "-extldflags -static" main.go

FROM alpine:3.12
COPY --from=builder /app/main /app/main
ENTRYPOINT [ "/app/main" ]
