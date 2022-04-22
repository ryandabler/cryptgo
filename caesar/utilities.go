package caesar

import (
	"cryptgo/internal/util"
	"sort"
)

func permutations(text []rune, alphabet []rune) [][]rune {
	ps := make([][]rune, len(alphabet))

	for i, v := range alphabet {
		ps[i] = decrypt(text, v, alphabet)
	}

	return ps
}

func toStrings(rs [][]rune) []string {
	ss := make([]string, len(rs))

	for i, v := range rs {
		ss[i] = string(v)
	}

	return ss
}

func sortedFreqs(freqs map[string]float64) []string {
	intermediate := make([]struct {
		k string
		v float64
	}, len(freqs))
	i := 0

	for k, v := range freqs {
		intermediate[i] = struct {
			k string
			v float64
		}{k, v}
		i++
	}

	sort.Slice(intermediate, func(i, j int) bool {
		return intermediate[j].v < intermediate[i].v
	})

	return util.MapSlice(intermediate, func(e struct {
		k string
		v float64
	}) string {
		return e.k
	})
}
