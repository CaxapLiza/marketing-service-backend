FROM golang:alpine as builder

WORKDIR /app/services/analytic/cmd

COPY /services/analytic/internal /app/services/analytic/internal
COPY /services/analytic/internal/handler /app/services/analytic/internal/handler
COPY /services/analytic/internal/repository /app/services/analytic/internal/repository

COPY /services/common /app/services/common

RUN go build -o main .
FROM alpine

WORKDIR /app/services/analytic/cmd

COPY --from=builder /app/services/analytic/cmd/main .

EXPOSE 7070

CMD ["./main"]
