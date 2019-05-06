package main

import (
	"log"
	"os"
	"time"

	pb "github.com/mattmoor/korpc-sample/gen/proto"
	"go.opencensus.io/trace"
	"golang.org/x/net/context"

	"github.com/mattmoor/korpc-sample/pkg/connection"
)

func main() {
	client := pb.NewOneClient(connection.Client())

	ctx := context.Background()

	ctx, span := trace.StartSpan(ctx, "client",
		// Always sample
		trace.WithSampler(trace.AlwaysSample()),
		// Sample 1% of requests
		// trace.WithSampler(trace.ProbabilitySampler(0.01)),
	)

	ctx, cancel := context.WithTimeout(ctx, 90*time.Second)
	defer cancel()

	start := time.Now()
	resp, err := client.Magic(ctx, &pb.Request{Msg: os.Args[1]})
	if err != nil {
		log.Fatalf("Magic() = %v", err)
	}

	log.Printf("Got back: %s", resp.GetMsg())
	log.Printf("took: %v (span: %v)", time.Since(start), span)
}
