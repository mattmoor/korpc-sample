package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	pb "github.com/mattmoor/korpc-sample/gen/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	serverAddr         = flag.String("server_addr", "127.0.0.1:8080", "The server address in the format of host:port")
	serverHostOverride = flag.String("server_host_override", "", "")
	insecure           = flag.Bool("insecure", false, "Set to true to skip SSL validation")
)

func main() {
	flag.Parse()

	var opts []grpc.DialOption
	if *serverHostOverride != "" {
		opts = append(opts, grpc.WithAuthority(*serverHostOverride))
	}
	if *insecure {
		opts = append(opts, grpc.WithInsecure())
	}

	conn, err := grpc.Dial(*serverAddr, opts...)
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()

	client := pb.NewSampleServiceClient(conn)

	unary(client, "hello")
	streamIn(client, "hello")
	streamOut(client, "hello")
	streamInOut(client, "hello")
}

func unary(client pb.SampleServiceClient, msg string) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	rep, err := client.Unary(ctx, &pb.Request{Msg: msg})
	if err != nil {
		log.Fatalf("Unary failed %v: ", err)
	}
	log.Printf("Unary got %v\n", rep.GetMsg())
}

func streamIn(client pb.SampleServiceClient, msg string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	stream, err := client.StreamIn(ctx)
	if err != nil {
		log.Fatalf("StreamIn =%v", err)
	}

	i := 0
	for i < 10 {
		if err := stream.Send(&pb.Request{Msg: fmt.Sprintf("%s-%d", msg, i)}); err != nil {
			log.Fatalf("Failed to send a stream message: %v", err)
		}
		i++
	}
	resp, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("StreamIn failed %v: ", err)
	}
	log.Printf("StreamIn got %v\n", resp.GetMsg())
}

func streamOut(client pb.SampleServiceClient, msg string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	stream, err := client.StreamOut(ctx, &pb.Request{Msg: msg})
	if err != nil {
		log.Fatalf("StreamOut =%v", err)
	}

	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a response : %v", err)
			}
			log.Printf("Got %s", in.GetMsg())
		}
	}()

	<-waitc
}

func streamInOut(client pb.SampleServiceClient, msg string) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	stream, err := client.Stream(ctx)
	if err != nil {
		log.Fatalf("Stream =%v", err)
	}

	waitc := make(chan struct{})
	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				// read done.
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Failed to receive a response : %v", err)
			}
			log.Printf("Got %s", in.GetMsg())
		}
	}()

	i := 0
	for i < 10 {
		if err := stream.Send(&pb.Request{Msg: fmt.Sprintf("%s-%d", msg, i)}); err != nil {
			log.Fatalf("Failed to send a stream message: %v", err)
		}
		i++
	}
	stream.CloseSend()
	<-waitc
}
