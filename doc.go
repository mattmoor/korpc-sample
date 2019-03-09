package sample

// These direct korpc to install the necessary tooling and
// generate all of the code and configuration necessary to
// develop a korpc service.

//go:generate korpc install
//go:generate korpc generate --base=github.com/mattmoor/korpc-sample service.proto
