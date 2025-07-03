package main

import (
	"context"
	"time"
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
