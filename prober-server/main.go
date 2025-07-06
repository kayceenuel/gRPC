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

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 50051, "the server port")
	// Define the counter metric variables
	ProbeRequestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "probe_request_total",
			Help: "Total number of probe requests",
		},
		[]string{"endpoint"},
	)

	probeLatencyMs = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "probe_Latency_miliseconds",
			Help: "Latest probe latency in miliseconds",
		},
		[]string{"endpoint"},
	)
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

	client := &http.Client{Timeout: 2 * time.Second}

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
	// Register metrics
	prometheus.MustRegister(ProbeRequestTotal)
	prometheus.MustRegister(probeLatencyMs)

	// Start HTTP server for metrics
	go func() {
		http.Handle("/metrics", promhttp.Handler())
		log.Println("Metrics server Listening on :8080")
		log.Fatal(http.ListenAndServe(":8080", nil))
	}()

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
