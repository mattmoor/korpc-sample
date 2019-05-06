package foo

import (
	"context"
	"crypto/sha1"
	"encoding/hex"

	pb "github.com/mattmoor/korpc-sample/gen/proto"
)

func Impl(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	h := sha1.New()
	h.Write([]byte(req.GetMsg()))
	sha1_hash := hex.EncodeToString(h.Sum(nil))
	return &pb.Response{Msg: sha1_hash}, nil
}
