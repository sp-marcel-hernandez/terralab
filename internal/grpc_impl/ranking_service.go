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

func (rs *RankingService) PutScore(ctx context.Context, req *pb.PlayerRecord) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, rs.Repository.SavePlayer(*toInternalPlayer(req))
}

func (rs *RankingService) GetScore(ctx context.Context, req *pb.GetScoreRequest) (*pb.Ranking, error) {
	ranking, err := rs.Repository.GetTopRanking(req.Top)

	if err != nil {
		return nil, err
	}

	res := &pb.Ranking{}
	for _, player := range ranking {
		res.Ranking = append(res.Ranking, toGrpcPlayer(&player))
	}

	return res, nil
}

func toInternalPlayer(p *pb.PlayerRecord) *common.Player {
	return &common.Player{Id: p.PlayerId, Score: p.Score}
}

func toGrpcPlayer(p *common.Player) *pb.PlayerRecord {
	return &pb.PlayerRecord{PlayerId: p.Id, Score: p.Score}
}
