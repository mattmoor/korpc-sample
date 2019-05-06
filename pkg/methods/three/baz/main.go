package baz

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	pb "github.com/mattmoor/korpc-sample/gen/proto"
)

func Impl(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	h := sha256.New()
	h.Write([]byte(req.GetMsg()))
	sha256_hash := hex.EncodeToString(h.Sum(nil))
	return &pb.Response{Msg: sha256_hash}, nil
}
