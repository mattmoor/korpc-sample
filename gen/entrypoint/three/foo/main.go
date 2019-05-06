package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/stats/view"

	pb "github.com/mattmoor/korpc-sample/gen/proto"
	impl "github.com/mattmoor/korpc-sample/pkg/methods/three/foo"
)

type server struct {
  pb.UnimplementedThreeServer
}


func (s *server) Foo(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return impl.Impl(ctx, req)
}


func main() {
	if len(os.Args) > 1 && os.Args[1] == "probe" {
		probe()
		return
	}

	// Prior to this, exporters should be registered via:
	// func init() {
	//    view.RegisterExporter(...)
	// }

	// Register the views to collect server request count.
	if err := view.Register(ocgrpc.DefaultServerViews...); err != nil {
		log.Fatal(err)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// The grpcServer is currently configured to serve h2c traffic by default.
	// To configure credentials or encryption, see: https://grpc.io/docs/guides/auth.html#go
	grpcServer := grpc.NewServer(grpc.StatsHandler(&ocgrpc.ServerHandler{}))

	pb.RegisterThreeServer(grpcServer, &server{})
	healthpb.RegisterHealthServer(grpcServer, &health{})

	grpcServer.Serve(lis)
}

// Based on github.com/grpc-ecosystem/grpc-health-probe
func probe() {
	ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
	defer cancel()
	conn, err := grpc.DialContext(ctx, fmt.Sprintf("localhost:%s", os.Getenv("PORT")), grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Error connecting: %v", err)
	}
	resp, err := healthpb.NewHealthClient(conn).Check(ctx, &healthpb.HealthCheckRequest{Service: "Three"})
	if err != nil {
		log.Fatalf("Error health checking: %v", err)
	}
	log.Printf("Health check: %#v", resp)
}

type health struct {}

func (h *health) Check(ctx context.Context, req *healthpb.HealthCheckRequest) (*healthpb.HealthCheckResponse, error) {
	return &healthpb.HealthCheckResponse{Status: healthpb.HealthCheckResponse_SERVING}, nil
}

func (h *health) Watch(*healthpb.HealthCheckRequest, healthpb.Health_WatchServer) error {
	return status.Error(codes.Unimplemented, "korpc does not implement Watch")
}

// Don't complain about the import
var _ = io.EOF
