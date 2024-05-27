FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

# Install golangci-lint
RUN apk add --no-cache curl git
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2

COPY . ./

WORKDIR /app/cmd/server

RUN go build -o /exchange-rate-receiver

CMD ["/exchange-rate-receiver"]
