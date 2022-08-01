# syntax=docker/dockerfile:1

FROM golang:1.18-alpine as builder

WORKDIR /app

# Download necessary Go modules
COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Copy source code and build
# COPY <src-path> <destination-path>
COPY . .
RUN go build -o golang-fiber main.go


# Deploy reduce size build
FROM alpine:3.10

WORKDIR /

COPY --from=builder /app/golang-fiber /

EXPOSE 3000

ENTRYPOINT ["/golang-fiber"]
