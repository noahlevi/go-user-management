FROM golang:1.23 AS builder

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

COPY .env ./.env 

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o user-service ./main.go

FROM alpine:latest  

WORKDIR /root/

COPY --from=builder /app/user-service .

EXPOSE 8080

CMD ["./user-service"]