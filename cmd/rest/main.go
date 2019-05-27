package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/golang/glog"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	gw "github.com/mattmoor/korpc-sample"
)

var (
	echoEndpoint = flag.String("echo_endpoint", "sample-service.default.svc.cluster.local:80", "endpoint of YourService")
)

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := gw.RegisterSampleServiceHandlerFromEndpoint(ctx, mux, *echoEndpoint, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(":"+os.Getenv("PORT"), mux)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
