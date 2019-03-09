package streamout

import (
	"context"
	"errors"

	pb "github.com/mattmoor/korpc-sample/gen/proto"
)


func Impl(ctx context.Context, req *pb.Request, resp chan *pb.Response) error {
	return errors.New(`You need to implement SampleService.StreamOut!!!`)
}

