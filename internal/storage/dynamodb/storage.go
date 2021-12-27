package dynamodb

import (
	"github.com/sp-marcel-hernandez/terralab/internal"
)

type DynamoStorage struct {
}

func (ms *DynamoStorage) SavePlayer(p internal.Player) error {
	return nil
}

func (ms *DynamoStorage) GetTopRanking(top uint64) (internal.Ranking, error) {
	return nil, nil
}
