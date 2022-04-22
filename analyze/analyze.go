package analyze

import (
	"cryptgo/internal/compset"
	"cryptgo/internal/util"
)

func BuildAlpha(text string) string {
	return compset.UniqueLetters(text)
}

func LetterFreq(text string, alphabet string) map[string]float64 {
	count := float64(len(text))
	alphaS := compset.FromString(alphabet)
	freqs := make(map[string]float64)

	for _, r := range text {
		c := string(r)

		if !alphaS.Has(c) {
			continue
		}

		freqs[c] += 1
	}

	return util.MapValues(freqs, func(v float64, k string) float64 { return v / count })
}
