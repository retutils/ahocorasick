// ahocorasick_test.go: test suite for ahocorasick
//
// Copyright (c) 2013 CloudFlare, Inc.

package ahocorasick

import (
	"testing"
)

func assert(t *testing.T, b bool) {
	if !b {
		t.Fail()
	}
}

func TestNoPatterns(t *testing.T) {
	m := NewStringMatcher([]string{})
	hits := m.Match([]byte("foo bar baz"))
	assert(t, len(hits) == 0)
}

func TestNoData(t *testing.T) {
	m := NewStringMatcher([]string{"foo", "baz", "bar"})
	hits := m.Match([]byte(""))
	assert(t, len(hits) == 0)
}

func TestSuffixes(t *testing.T) {
	m := NewStringMatcher([]string{"Superman", "uperman", "perman", "erman"})
	hits := m.Match([]byte("The Man Of Steel: Superman"))
	assert(t, len(hits) == 4)
}

func TestPrefixes(t *testing.T) {
	m := NewStringMatcher([]string{"Superman", "Superma", "Superm", "Super"})
	hits := m.Match([]byte("The Man Of Steel: Superman"))
	assert(t, len(hits) == 4)
}

func TestOffset(t *testing.T) {
	m := NewStringMatcher([]string{"Superman", "Man", "Of", "The", "Steel"})
	hits := m.Match([]byte("The Man Of Steel: Superman"))
	assert(t, len(hits) == 5)
	assert(t, hits["Superman"] == 25)
	assert(t, hits["The"] == 2)
	assert(t, hits["Steel"] == 15)
	assert(t, hits["Man"] == 6)
	assert(t, hits["Of"] == 9)
}
