FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o service .

CMD ["./main"]
