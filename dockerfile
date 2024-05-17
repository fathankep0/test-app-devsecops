FROM golang:alpine AS builder

WORKDIR /app

COPY main.go .

RUN go mod init test

RUN go build -o test .

FROM alpine

RUN apk --no-cache add ca-certificates

WORKDIR /app

COPY --from=builder /app/test .

# Menambahkan environment variable
ENV Environment_Name=Poduction

CMD ["./test"]
