FROM golang:1.21.4-alpine AS builder

WORKDIR /build

COPY . .

RUN go build -ldflags="-s -w" -o  main main.go

FROM alpine

WORKDIR /build

COPY --from=builder /build/main /build/main

CMD ["./main"]