package sample

// These direct korpc to install the necessary tooling and
// generate all of the code and configuration necessary to
// develop a korpc service.

//    go:generate korpc install
//    go:generate korpc generate --base=github.com/mattmoor/korpc-sample --domain=api.mattmoor.io service.proto

//go:generate /usr/local/google/home/mattmoor/go/bin/.korpc/protoc-3.7.0/bin/protoc -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis  --go_out=plugins=grpc:. service.proto
//go:generate /usr/local/google/home/mattmoor/go/bin/.korpc/protoc-3.7.0/bin/protoc -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis  --grpc-gateway_out=logtostderr=true:. service.proto
