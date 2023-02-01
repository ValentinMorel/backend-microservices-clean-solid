FROM golang:1.19-alpine AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go mod download
WORKDIR /build/cmd/
RUN go build -o gateway-service

FROM golang:1.19-alpine
COPY --from=builder /build/cmd/gateway-service /app/
WORKDIR /app
CMD ["./gateway-service"]