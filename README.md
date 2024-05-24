## Protobuf
### Install
(MacOS)
To install `protoc` run
```bash
brew install protobuf
```
Then download Go protobuf plugins
```bash
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
### Generation
Navigate to the root directory of your project and run the following commands to generate the Go code from the protobuf definitions:

```bash
protoc --go_out=.. --go-grpc_out=.. proto/exchange_rate.proto
```

