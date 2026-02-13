FROM     golang:1.21-alpine AS builder
WORKDIR  /app
COPY     go.mod go.sum ./
RUN      go mod download
COPY     . .
RUN      go build -o main .

FROM     alpine:latest
WORKDIR  /root/
COPY     --from=builder /app/main .
RUN      mkdir uploads
EXPOSE   3000
CMD      ["./main"]