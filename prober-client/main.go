package main

import (
	"flag"
	"log"

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
	cleint := pb.NewProberClient(conn)
}
