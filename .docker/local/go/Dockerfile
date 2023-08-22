FROM golang:1.20.4-alpine

RUN apk update && \
    apk add --no-cache git make gcc libc-dev && \
    go install github.com/swaggo/swag/cmd/swag@v1.8.3 && \
    go install github.com/go-delve/delve/cmd/dlv@latest && \
    go install github.com/cosmtrek/air@latest
