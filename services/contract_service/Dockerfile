FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o contract_service .

CMD ["./main"]
