package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	pb "grpc-prober/prober"

	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "the server port")
)

// the server is used to implment the proberService
type server struct {
	pb.UnimplementedProberServiceServer
}

func (s *server) DoProbes(ctx context.Context, in *pb.ProbeRequest) (*pb.ProbeReply, error) {
	endpoint := in.GetEndpoint()
	repetitions := in.GetRepetitions()

	var total time.Duration
	var successCount int

	client := &http.Client{Timeout: 5 * time.Second}

	for i := 0; i < int(repetitions); i++ {
		start := time.Now()
		resp, err := client.Get(endpoint)
		if err != nil {
			log.Printf("Request %d failed: %v", i+1, err)
			continue
		}
		resp.Body.Close()
		elapsed := time.Since(start)

		total += elapsed
		successCount++
	}

	var avgMs float32
	if successCount > 0 {
		avgMs = float32(total.Milliseconds()) / float32(successCount)
	}

	return &pb.ProbeReply{
		LatencyMsecs: avgMs,
	}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterProberServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v", err)
	}
}
