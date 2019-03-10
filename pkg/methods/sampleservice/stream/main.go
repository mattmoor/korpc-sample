package stream

import (
	"context"
	"fmt"
	"os"
	"time"

	pb "github.com/mattmoor/korpc-sample/gen/proto"
)

func Impl(ctx context.Context, req <-chan *pb.Request, resp chan *pb.Response) error {
	for {
		select {
		case _, ok := <-req:
			if !ok {
				return nil
			}
			resp <- &pb.Response{
				Msg: fmt.Sprintf("pong %s %s", time.Now(), os.Getenv("WHOAMI")),
			}
		}
	}
}
