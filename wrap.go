package fzflib

import (
	"github.com/sshelll/fzflib/algo"
	"github.com/sshelll/fzflib/util"
)

type Fzf struct {
	targets       []string
	casesensitive bool
	normalize     bool
	forward       bool
	withPos       bool
}

func New() *Fzf {
	// only set normalize as true
	return &Fzf{
		normalize: true,
	}
}

func (f *Fzf) CaseSensitive(sensitivity bool) *Fzf {
	f.casesensitive = sensitivity
	return f
}

func (f *Fzf) Normalize(normalize bool) *Fzf {
	f.normalize = normalize
	return f
}

func (f *Fzf) Forward(forward bool) *Fzf {
	f.forward = forward
	return f
}

func (f *Fzf) WithPos(withPos bool) *Fzf {
	f.withPos = withPos
	return f
}

func (f *Fzf) AppendTargets(inputs ...string) *Fzf {
	f.targets = append(f.targets, inputs...)
	return f
}

func (f *Fzf) Clear() *Fzf {
	f.targets = []string{}
	return f
}

func (f *Fzf) Match(pattern string) []*MatchResult {
	results := []*MatchResult{}
	for i := range f.targets {
		t := f.targets[i]
		chars := util.ToChars([]byte(t))
		r, pos := algo.FuzzyMatchV2(
			f.casesensitive,
			f.normalize,
			true,
			&chars,
			[]rune(pattern),
			true,
			nil,
		)
		if r.Score == 0 {
			continue
		}
		results = append(results, &MatchResult{
			content: &t,
			score:   r.Score,
			pos:     pos,
		})
	}
	return results
}
