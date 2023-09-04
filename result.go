package fzflib

type MatchResult struct {
	content *string
	score   int
	pos     *[]int
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
