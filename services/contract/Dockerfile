FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o contract .

CMD ["./main"]
