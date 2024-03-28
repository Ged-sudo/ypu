FROM golang:1.23 AS builder

RUN apk --no-cahe add bash git make gcc gettext

WORKDIR /usr/local/src

# .COPY ["go.mod", "go.sum"]
RUN go mod download