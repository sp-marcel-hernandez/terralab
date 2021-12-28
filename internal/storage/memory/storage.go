package memory

import (
	"sort"

	common "github.com/sp-marcel-hernandez/terralab/internal"
)

type memoryStorage struct {
	storage map[string]common.Player
}

func NewMemoryStorage() *memoryStorage {
	ms := &memoryStorage{}
	ms.storage = make(map[string]common.Player)

	return ms
}

func (ms *memoryStorage) SavePlayer(p common.Player) error {
	ms.storage[p.Id] = p

	return nil
}

func (ms *memoryStorage) GetTopRanking(top uint64) (common.Ranking, error) {
	var ranking common.Ranking

	for _, player := range ms.storage {
		ranking = append(ranking, player)
	}

	sort.Sort(byScore(ranking))

	if top >= uint64(len(ranking)) {
		return ranking, nil
	}

	return ranking[:top], nil
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
