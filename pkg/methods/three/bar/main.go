package bar

import (
	"context"
	"crypto/md5"
	"encoding/hex"

	pb "github.com/mattmoor/korpc-sample/gen/proto"
)

func Impl(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	h := md5.New()
	h.Write([]byte(req.GetMsg()))
	md5_hash := hex.EncodeToString(h.Sum(nil))
	return &pb.Response{Msg: md5_hash}, nil
}
