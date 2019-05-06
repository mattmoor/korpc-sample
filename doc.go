package sample

// These direct korpc to install the necessary tooling and
// generate all of the code and configuration necessary to
// develop a korpc service.

//go:generate go install github.com/mattmoor/korpc/cmd/korpc
//go:generate korpc install
//go:generate korpc generate --base=github.com/mattmoor/korpc-sample --domain=api.mattmoor.io service.proto
