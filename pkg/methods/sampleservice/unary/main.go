package unary

import (
	"context"
	"errors"

	pb "github.com/mattmoor/korpc-sample/gen/proto"
)


func Impl(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	return nil, errors.New(`You need to implement SampleService.Unary!!!`)
}

