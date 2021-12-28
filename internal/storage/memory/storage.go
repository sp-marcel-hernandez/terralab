package memory

import (
	"sort"
	"sync"

	common "github.com/sp-marcel-hernandez/terralab/internal"
)

type memoryStorage struct {
	mux     sync.Mutex
	storage map[string]common.Player
}

func NewMemoryStorage() *memoryStorage {
	ms := &memoryStorage{}
	ms.storage = make(map[string]common.Player)

	return ms
}

func (ms *memoryStorage) SavePlayer(p common.Player) error {
	ms.mux.Lock()
	defer ms.mux.Unlock()

	ms.storage[p.Id] = p

	return nil
}

func (ms *memoryStorage) GetTopRanking(top uint64) (common.Ranking, error) {
	ms.mux.Lock()
	defer ms.mux.Unlock()

	ranking := make(common.Ranking, len(ms.storage))

	i := 0
	for _, player := range ms.storage {
		ranking[i] = player
		i++
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
