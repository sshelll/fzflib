package fzflib

import (
	"unicode"

	"github.com/sshelll/fzflib/algo"
	"github.com/sshelll/fzflib/util"
)

type Item struct {
	Content string
	Any     any
}

type Fzf struct {
	targets       []string
	items         []*Item
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

func (f *Fzf) AppendItems(items ...*Item) *Fzf {
	f.items = append(f.items, items...)
	return f
}

func (f *Fzf) Clear() *Fzf {
	f.targets = []string{}
	return f
}

// Match matches the pattern with targets(string slice).
func (f *Fzf) Match(pattern string) []*MatchResult {
	results := f.match(pattern, f.casesensitive)
	return results
}

// MergeMatch merges the case-sensitive and case-insensitive results.
func (f *Fzf) MergeMatch(pattern string) []*MatchResult {
	results := f.mergeMatch(pattern, f.match)
	return results
}

// MatchItem matches the pattern with items, see Item struct.
func (f *Fzf) MatchItem(pattern string) []*MatchResult {
	results := f.matchItem(pattern, f.casesensitive)
	return results
}

// MergeMatchItem merges the case-sensitive and case-insensitive results.
func (f *Fzf) MergeMatchItem(pattern string) []*MatchResult {
	results := f.mergeMatch(pattern, f.matchItem)
	return results
}

func (f *Fzf) mergeMatch(pattern string, matchFn func(string, bool) []*MatchResult) []*MatchResult {
	hasUpper := false
	for _, c := range pattern {
		if unicode.IsUpper(c) {
			hasUpper = true
			break
		}
	}
	r1 := matchFn(pattern, true)
	r2 := matchFn(pattern, false)
	rset := make(map[string]*MatchResult, len(r2))
	for i := range r1 {
		r := r1[i]
		rset[r.Content()] = r
	}
	for i := range r2 {
		r := r2[i]
		if _, ok := rset[r.Content()]; !ok {
			rset[r.Content()] = r
		} else if hasUpper {
			// only add two scores when users input upper chars
			rset[r.Content()].score += r.score
		}
	}
	results := make([]*MatchResult, 0, len(rset))
	for k := range rset {
		results = append(results, rset[k])
	}
	return results
}

func (f *Fzf) match(pattern string, csensitive bool) []*MatchResult {
	results := []*MatchResult{}
	for i := range f.targets {
		t := f.targets[i]
		chars := util.ToChars([]byte(t))
		r, pos := algo.FuzzyMatchV2(
			csensitive,
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

func (f *Fzf) matchItem(pattern string, csensitive bool) []*MatchResult {
	results := []*MatchResult{}
	for i := range f.items {
		t := f.items[i]
		chars := util.ToChars([]byte(t.Content))
		r, pos := algo.FuzzyMatchV2(
			csensitive,
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
			content: &t.Content,
			score:   r.Score,
			pos:     pos,
			item:    t,
		})
	}
	return results
}
