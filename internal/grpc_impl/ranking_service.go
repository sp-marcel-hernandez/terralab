package grpc_impl

import (
	"context"

	common "github.com/sp-marcel-hernandez/terralab/internal"
	"github.com/sp-marcel-hernandez/terralab/pkg/pb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RankingService struct {
	pb.UnimplementedRankingServiceServer
	Repository common.RankingRepository
}

func (s *RankingService) PutScore(ctx context.Context, p *pb.PlayerRecord) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, s.Repository.SavePlayer(*toInternalPlayer(p))
}

func (s *RankingService) GetScore(ctx context.Context, gp *pb.GetScoreRequest) (*pb.Ranking, error) {
	ir, err := s.Repository.GetTopRanking(gp.Top)

	if err != nil {
		return nil, err
	}

	gr := &pb.Ranking{}
	for _, p := range ir {
		gr.Ranking = append(gr.Ranking, toGrpcPlayer(&p))
	}

	return gr, nil
}

func toInternalPlayer(p *pb.PlayerRecord) *common.Player {
	return &common.Player{Id: p.PlayerId, Score: p.Score}
}

func toGrpcPlayer(p *common.Player) *pb.PlayerRecord {
	return &pb.PlayerRecord{PlayerId: p.Id, Score: p.Score}
}
