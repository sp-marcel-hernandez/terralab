package grpc_impl

import (
	"context"

	"github.com/sp-marcel-hernandez/terralab/pkg/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RankingService struct {
	pb.UnimplementedRankingServiceServer
}

func (s *RankingService) PutScore(context.Context, *pb.PlayerRecord) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}

func (s *RankingService) GetScore(context.Context, *pb.GetScoreRequest) (*pb.Ranking, error) {
	return &pb.Ranking{Ranking: []*pb.PlayerRecord{
		{PlayerId: "100", Score: 1000},
		{PlayerId: "101", Score: 999},
	}}, nil
}
