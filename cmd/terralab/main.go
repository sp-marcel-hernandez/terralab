package main

import (
	"log"
	"net"
	"os"

	"github.com/sp-marcel-hernandez/terralab/internal/grpc_impl"
	"github.com/sp-marcel-hernandez/terralab/pkg/pb"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", os.Getenv("SP_LISTEN_ADDR"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", listener.Addr())

	s := grpc.NewServer()
	pb.RegisterRankingServiceServer(s, &grpc_impl.RankingService{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
