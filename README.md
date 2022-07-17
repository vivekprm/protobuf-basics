# protobuf-basics

# Prerequisites
- Install protoc compiler from https://github.com/protocolbuffers/protobuf/releases.
- Install protogen plugin for go.
```sh
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```

# Generate protobuf files.
```sh
protoc -I=. --go_out=pbgo --go_opt=paths=source_relative ./addressbook.proto
```

# To run the app
```sh
go run main.go
```
