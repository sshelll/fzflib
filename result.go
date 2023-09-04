package fzflib

type MatchResult struct {
	content *string
	score   int
}

func (r *MatchResult) Content() string {
	return *r.content
}

func (r *MatchResult) Score() int {
	return r.score
}
