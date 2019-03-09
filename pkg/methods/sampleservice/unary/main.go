package unary

import (
	"context"
	"fmt"

	pb "github.com/mattmoor/korpc-sample/gen/proto"
)

func Impl(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return &pb.Response{Msg: fmt.Sprintf("%s - pong", req.Msg)}, nil
}
