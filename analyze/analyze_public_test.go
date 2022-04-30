package analyze_test

import (
	"cryptgo/analyze"
	"sort"
	"testing"
)

const text = "bench"
const alpha = "bench"
const duplicateText = "Hello world"
const duplicateAlpha = "Helo wrd"

type sortableRunes []rune

func (s sortableRunes) Swap(i int, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortableRunes) Less(i int, j int) bool {
	return s[i] < s[j]
}

func (s sortableRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	runes := sortableRunes([]rune(s))
	sort.Sort(runes)
	return string(runes)
}

func TestBuildAlpha(t *testing.T) {
	// Case: build alphabet from string
	result := analyze.BuildAlpha(text)
	if sortString(result) != sortString(alpha) {
		t.Errorf("Case alphabet: From text %q expected alphabet %q -- got %q", text, alpha, result)
	}

	// Case: alphabet should only have unique letters
	result = analyze.BuildAlpha(duplicateText)
	if sortString(result) != sortString(duplicateAlpha) {
		t.Errorf("Case de-dupe: From text %q expected alphabet %q -- got %q", duplicateText, duplicateAlpha, result)
	}
}
