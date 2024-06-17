# Stage 1: Build stage
FROM golang:1.18 AS builder

LABEL maintainer="Anjush Bhargavan <anjushbhargavan12@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/app main.go

# Stage 2: Final stage
FROM alpine:latest

LABEL maintainer="Anjush Bhargavan <anjushbhargavan12@gmail.com>"

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/bin/app .

COPY --from=builder /app/config ./config

COPY --from=builder /app/docs ./docs

EXPOSE 8080

CMD ["./app"]
