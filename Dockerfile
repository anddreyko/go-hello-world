FROM golang:1.21.4-alpine AS builder

WORKDIR /build

COPY . .

RUN cd ./src/hello && go build -ldflags="-s -w" -o  hello hello.go

FROM alpine

WORKDIR /app

COPY --from=builder /build/src/hello/hello /app/main

CMD ["./main"]
