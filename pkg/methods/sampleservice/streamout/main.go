package streamout

import (
	"context"
	"fmt"
	"os"

	pb "github.com/mattmoor/korpc-sample/gen/proto"
)

func Impl(ctx context.Context, req *pb.Request, resp chan *pb.Response) error {
	for i := 0; i < 10; i++ {
		resp <- &pb.Response{
			Msg: fmt.Sprintf("%s %d/10 (%s)", req.Msg, i, os.Getenv("WHOAMI")),
		}
	}
	return nil
}
