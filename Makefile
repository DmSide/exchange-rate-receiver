.PHONY: build
build:
	cd cmd/server && go build -o ../../exchange_rate-receiver && cd ../..

.PHONY: test
test:
	go test ./...

.PHONY: docker-build
docker-build:
	docker build -t exchange-rate-receiver .

.PHONY: run
run:
	go run ./cmd/server

.PHONY: lint
lint:
	golangci-lint run