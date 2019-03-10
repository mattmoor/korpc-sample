package unary

import (
	"context"
	"fmt"
	"os"

	pb "github.com/mattmoor/korpc-sample/gen/proto"
)

func Impl(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{Msg: fmt.Sprintf("%s - pong - %s", req.Msg, os.Getenv("WHOAMI"))}, nil
}
