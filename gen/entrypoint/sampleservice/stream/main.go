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

	pb "github.com/mattmoor/korpc-sample/gen/proto"
	impl "github.com/mattmoor/korpc-sample/pkg/methods/sampleservice/stream"
)

type server struct {
  pb.UnimplementedSampleServiceServer
}


func (s *server) Stream(stream pb.SampleService_StreamServer) error {
	input := make(chan *pb.Request)
	output := make(chan *pb.Response)

	errCh := make(chan error)

	go func() {
		defer close(input)
		for {
			req, err := stream.Recv()
			if err == io.EOF {
				return
			}
			if err != nil {
				errCh <- err
				return
			}
			input <- req
		}
	}()

	go func() {
		defer close(output)
		errCh <- impl.Impl(stream.Context(), input, output)
	}()

	go func() {
		defer close(errCh)
		for {
			select {
			case resp, ok := <-output:
				if !ok {
					// Quit when the output channel closes.
					return
				}
				if err := stream.Send(resp); err != nil {
					errCh <- err
					return
				}
			}
		}
	}()

	return <-errCh
}


func main() {
	if len(os.Args) > 1 && os.Args[1] == "probe" {
		probe()
		return
	}
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// The grpcServer is currently configured to serve h2c traffic by default.
	// To configure credentials or encryption, see: https://grpc.io/docs/guides/auth.html#go
	grpcServer := grpc.NewServer()

	pb.RegisterSampleServiceServer(grpcServer, &server{})
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
	resp, err := healthpb.NewHealthClient(conn).Check(ctx, &healthpb.HealthCheckRequest{Service: "SampleService"})
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
