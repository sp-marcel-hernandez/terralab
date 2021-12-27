package grpc_impl

import (
	"context"

	"github.com/sp-marcel-hernandez/terralab/internal"
	"github.com/sp-marcel-hernandez/terralab/pkg/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RankingService struct {
	pb.UnimplementedRankingServiceServer
	Repository internal.RankingRepository
}

func (s *RankingService) PutScore(ctx context.Context, gp *pb.PlayerRecord) (*emptypb.Empty, error) {
	ip := &internal.Player{
		Id:    gp.PlayerId,
		Score: gp.Score,
	}

	return &emptypb.Empty{}, s.Repository.SavePlayer(*ip)
}

func (s *RankingService) GetScore(ctx context.Context, gp *pb.GetScoreRequest) (*pb.Ranking, error) {
	ir, err := s.Repository.GetTopRanking(gp.Top)

	if err != nil {
		return nil, err
	}

	gr := &pb.Ranking{}
	for _, ip := range ir {
		gr.Ranking = append(gr.Ranking, &pb.PlayerRecord{PlayerId: ip.Id, Score: ip.Score})
	}

	return gr, nil
}
