package dynamodb

import common "github.com/sp-marcel-hernandez/terralab/internal"

type DynamoStorage struct {
}

func (ds *DynamoStorage) SavePlayer(p common.Player) error {
	return nil
}

func (ds *DynamoStorage) GetTopRanking(top uint64) (common.Ranking, error) {
	return nil, nil
}
