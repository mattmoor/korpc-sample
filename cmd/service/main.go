package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	pb "github.com/mattmoor/korpc-sample"
)

type server struct{}

func (s *server) Unary(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{Msg: "Gotcha: " + req.GetMsg()}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", os.Getenv("PORT")))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// The grpcServer is currently configured to serve h2c traffic by default.
	// To configure credentials or encryption, see: https://grpc.io/docs/guides/auth.html#go
	grpcServer := grpc.NewServer()

	pb.RegisterSampleServiceServer(grpcServer, &server{})

	grpcServer.Serve(lis)
}
