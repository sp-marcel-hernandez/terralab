package main

import (
	"context"
	"github.com/sp-marcel-hernandez/terralab/internal/pb"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"net"
	"os"
)

type rankingService struct {
	pb.UnimplementedRankingServiceServer
}

func (s *rankingService) PutScore(context.Context, *pb.PlayerRecord) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *rankingService) GetScore(context.Context, *pb.GetScoreRequest) (*pb.Ranking, error) {
	return &pb.Ranking{Ranking: []*pb.PlayerRecord{
		{PlayerId: "100", Score: 1000},
		{PlayerId: "101", Score: 999},
	}}, nil
}

func main() {
	listener, err := net.Listen("tcp", os.Getenv("SP_LISTEN_ADDR"))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("server listening at %v", listener.Addr())

	s := grpc.NewServer()
	pb.RegisterRankingServiceServer(s, &rankingService{})

	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
