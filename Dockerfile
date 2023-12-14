FROM golang:latest
COPY go.mod .
COPY go.sum .
RUN go mod download

FROM golang:latest as analytic-builder
WORKDIR /app/analytic
COPY services/analytic .
RUN go build -o analytic .

FROM golang:latest as client-builder
WORKDIR /app/client
COPY services/client .
RUN go build -o client .

FROM golang:latest as contract-builder
WORKDIR /app/contract
COPY services/contract .
RUN go build -o contract .

FROM golang:latest as contract_service-builder
WORKDIR /app/contract_service
COPY services/contract_service .
RUN go build -o contract_service .

FROM golang:latest as project-builder
WORKDIR /app/project
COPY services/project .
RUN go build -o project .

FROM golang:latest as resource-builder
WORKDIR /app/resource
COPY services/resource .
RUN go build -o resource .

FROM golang:latest as service-builder
WORKDIR /app/service
COPY services/service .
RUN go build -o service .

WORKDIR /app

COPY --from=analytic-builder /app/analytic/analytic ./analytic
COPY --from=client-builder /app/client/client ./client
COPY --from=contract-builder /app/contract/contract ./contract
COPY --from=contract_service-builder /app/contract_service/contract_service ./contract_service
COPY --from=project-builder /app/project/project ./project
COPY --from=resource-builder /app/resource/resource ./resource
COPY --from=service-builder /app/service/service ./service

CMD ["./analytic", "./client", "./contract", "./contract_service", "./project", "./resource", "./service"]
