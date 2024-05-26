FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

WORKDIR /app/cmd/server

RUN go build -o /exchange-rate-receiver

CMD ["/exchange-rate-receiver"]
