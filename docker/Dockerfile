############## BUILD ##############

# syntax=docker/dockerfile:1

FROM golang:1.24 AS builder
WORKDIR /app

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64


# deps 캐시
COPY go.mod go.sum ./
RUN go mod download

# 앱 소스
COPY . .

RUN go build -trimpath -ldflags="-s -w" -o /probe-test-build .

############## RUN ##############

FROM scratch

WORKDIR /

COPY --from=builder /probe-test-build /probe-test

ENTRYPOINT ["/probe-test"]
