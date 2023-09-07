package fzflib

import "testing"

func TestMatch(t *testing.T) {
	fzf := New().CaseSensitive(false).Normalize(true)
	fzf.AppendTargets("foo", "bar", "baz", "FOOBaz")
	for _, r := range fzf.Match("o") {
		t.Log(r.Content(), r.Score(), r.Pos())
	}
	t.Log("----")
	for _, r := range fzf.Match("ba") {
		t.Log(r.Content(), r.Score(), r.Pos())
	}
	t.Log("----")
	for _, r := range fzf.Match("z") {
		t.Log(r.Content(), r.Score(), r.Pos())
	}
	t.Log("----")
	for _, r := range fzf.Match("F") {
		t.Log(r.Content(), r.Score(), r.Pos())
	}
	t.Log("----")
	for _, r := range fzf.MergeMatch("F") {
		t.Log(r.Content(), r.Score(), r.Pos())
	}
}

func TestMatchItem(t *testing.T) {
	fzf := New().CaseSensitive(false).Normalize(true)
	fzf.AppendItems(
		&Item{Content: "foo", Any: "foo_any"},
		&Item{Content: "bar", Any: 2},
		&Item{Content: "baz", Any: 3.14},
		&Item{Content: "FOOBaz", Any: "FOOBaz_any"},
	)
	for _, r := range fzf.MatchItem("a") {
		item := r.Item()
		t.Log(item.Content, item.Any, r.Score(), r.Pos())
	}
	t.Log("----")
	for _, r := range fzf.MergeMatchItem("F") {
		t.Log(r.Content(), r.Score(), r.Pos())
	}
}
