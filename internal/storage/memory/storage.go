package memory

import (
	"sort"

	common "github.com/sp-marcel-hernandez/terralab/internal"
)

type MemoryStorage struct {
	inMemoryRanking common.Ranking
}

func (ms *MemoryStorage) SavePlayer(p common.Player) error {
	ms.inMemoryRanking = append(ms.inMemoryRanking, p)
	sort.Sort(byScore(ms.inMemoryRanking))

	return nil
}

func (ms *MemoryStorage) GetTopRanking(top uint64) (common.Ranking, error) {
	if top >= uint64(len(ms.inMemoryRanking)) {
		return ms.inMemoryRanking, nil
	}

	return ms.inMemoryRanking[:top], nil
}

type byScore common.Ranking

func (bs byScore) Len() int {
	return len(bs)
}

func (bs byScore) Swap(i, j int) {
	bs[i], bs[j] = bs[j], bs[i]
}

func (bs byScore) Less(i, j int) bool {
	return bs[i].Score > bs[j].Score
}
