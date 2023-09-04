package fzflib

import "testing"

func TestWrap(t *testing.T) {
	fzf := New().CaseSensitive(false).Normalize(true)
	fzf.AppendTargets("foo", "bar", "baz")
	for _, r := range fzf.Match("o") {
		t.Log(r.Content(), r.Score())
	}
	t.Log("----")
	for _, r := range fzf.Match("ba") {
		t.Log(r.Content(), r.Score())
	}
	t.Log("----")
	for _, r := range fzf.Match("z") {
		t.Log(r.Content(), r.Score())
	}
}
