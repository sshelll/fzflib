package fzflib

type MatchResult struct {
	content *string
	score   int
	pos     *[]int
	item    *Item
}

func (r *MatchResult) Content() string {
	return *r.content
}

func (r *MatchResult) Score() int {
	return r.score
}

func (r *MatchResult) Pos() []int {
	if r.pos == nil {
		return []int{}
	}
	return *r.pos
}

func (r *MatchResult) Item() *Item {
	return r.item
}
