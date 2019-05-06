package connection

import (
	"log"

	"cloud.google.com/go/compute/metadata"
	"contrib.go.opencensus.io/exporter/stackdriver"
	"go.opencensus.io/plugin/ocgrpc"
	"go.opencensus.io/trace"
	"google.golang.org/grpc"
)

var client *grpc.ClientConn

func init() {
	projectID, err := metadata.ProjectID()
	if err != nil {
		log.Fatalf("Unable to fetch GCP ProjectID: %v", err)
	}

	// Create and register a OpenCensus Stackdriver Trace exporter.
	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID: projectID,
	})
	if err != nil {
		log.Fatalf("stackdriver.NewExporter() = %v", err)
	}
	trace.RegisterExporter(exporter)

	conn, err := grpc.Dial("api.mattmoor.io:80",
		grpc.WithInsecure(),
		grpc.WithStatsHandler(&ocgrpc.ClientHandler{}),
	)
	if err != nil {
		log.Fatalf("grpc.Dial() = %v", err)
	}
	client = conn
}

func Client() *grpc.ClientConn {
	return client
}
