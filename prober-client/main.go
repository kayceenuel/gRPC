package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	serverAddr = flag.String("serverAddr", "localhost:50051", "the address to connec to")
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(*serverAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewProberClient(conn)

	//Unary call
	cxt, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer.cancel()

	res, err := client.DoProbes(ctx, &pb.ProbeRequest{
		in.GetEndpoint(),
		in.GetRepetitions(),
	})
	if err != nil {
		fmt.Printf("Average response time: %f", res.GetLatencyMsecs())
	}
}
