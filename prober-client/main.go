package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "grpc-prober/prober"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr = flag.String("serverAddr", "localhost:50051", "the address to connec to")
)

func main() {
	flag.Parse()
	// set up a connection to the server.
	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewProberServiceClient(conn)

	// connect the server and print out the response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	res, err := client.DoProbes(ctx, &pb.ProbeRequest{
		Endpoint:    "http://www.google.com",
		Repetitions: 2,
	})
	if err != nil {
		log.Fatalf("probe failed: %v", err) // Exit on error
	}
	// Print response only if there's no error
	fmt.Printf("Average response time: %f ms\n", res.GetLatencyMsecs())
}
