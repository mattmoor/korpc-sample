package streamin

import (
	"context"
	"errors"

	pb "github.com/mattmoor/korpc-sample/gen/proto"
)


func Impl(ctx context.Context, req <-chan *pb.Request) (*pb.Response, error) {
	for {
		select {
		case _, ok := <-req:
			if !ok {
				return nil, errors.New("You need to implement SampleService.StreamIn!!!")
			}
		}
	}

}

