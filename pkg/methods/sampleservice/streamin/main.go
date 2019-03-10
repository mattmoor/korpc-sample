package streamin

import (
	"context"
	"fmt"
	"os"

	pb "github.com/mattmoor/korpc-sample/gen/proto"
)

func Impl(ctx context.Context, req <-chan *pb.Request) (*pb.Response, error) {
	resp := &pb.Response{}
	count := 0
	for {
		select {
		case req, ok := <-req:
			if !ok {
				return resp, nil
			}
			count++
			resp.Msg = fmt.Sprintf("%d received, last: %s (%s)", count, req.Msg, os.Getenv("WHOAMI"))
		}
	}
}
