package blah

import (
	"fmt"

	pb "github.com/mattmoor/korpc-sample/gen/proto"
	"github.com/mattmoor/korpc-sample/pkg/connection"
	"golang.org/x/net/context"
)

func Impl(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	client := pb.NewThreeClient(connection.Client())

	fooResp, err := client.Foo(ctx, &pb.Request{Msg: req.GetMsg()})
	if err != nil {
		return nil, err
	}

	return &pb.Response{Msg: fmt.Sprintf("blur: %q", fooResp.GetMsg())}, nil
}
