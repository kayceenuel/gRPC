package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

var (
	port = flag.int("port", 50051, "the server port")
)

// the server is used to implment the proberService
type server struct {
	pb.UnImplementedProberServer
}

func (s *server) DoProbes(ctx context.Context, in *pb.ProRequest) (*pb.ProbeReply, error) {
	start = time.Now()
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProberServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
