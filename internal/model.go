package common

type Player struct {
	Id    string
	Score uint64
}

// Players must be sorted by Score in descending order
type Ranking []Player

type RankingRepository interface {
	SavePlayer(p Player) error
	GetTopRanking(top uint64) (Ranking, error)
}
