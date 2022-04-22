package caesar

import "cryptgo/internal/util"

type Ranker func([]string) []string

func encrypt(text []rune, key rune, alphabet []rune) []rune {
	runes := []rune(text)
	mod := len(alphabet)
	shift := util.IndexOf(alphabet, key)
	cipher := make([]rune, len(runes))

	for i, r := range runes {
		pos := util.IndexOf(alphabet, r)
		newPos := util.ShiftWrap(pos, shift, mod, util.Left)
		cipher[i] = alphabet[newPos]
	}

	return cipher
}

func decrypt(text []rune, key rune, alphabet []rune) []rune {
	runes := []rune(text)
	mod := len(alphabet)
	shift := util.IndexOf(alphabet, key)
	plain := make([]rune, len(runes))

	for i, r := range runes {
		pos := util.IndexOf(alphabet, r)
		newPos := util.ShiftWrap(pos, shift, mod, util.Right)
		plain[i] = alphabet[newPos]
	}

	return plain
}

// Encrypt a string `text` of alphabet `alphabet` shifting left
// by a constant `key`.
//
// Each character of `text` will be left-shifted by the
// numerical value of `key` as it exists in the alphabet.
// This means the magnitude of left-shifting is based on
// where `k` appears in the alphabet. Thus if we shift
// by 'b' in alphabet 'abc' we would shift by 1 whereas
// in the alphabet 'acb' we would shift by 2.
func Encrypt(text string, key string, alphabet string) string {
	k := []rune(key)[0]
	cipher := encrypt([]rune(text), k, []rune(alphabet))
	return string(cipher)
}

// Decrypt a string `text` of alphabet `alphabet` shifting right
// by a constant `key`.
//
// Each character of `text` will be right-shifted by the
// numerical value of `key` as it exists in the alphabet.
// This means the magnitude of right-shifting is based on
// where `k` appears in the alphabet. Thus if we shift
// by 'b' in alphabet 'abc' we would shift by 1 whereas
// in the alphabet 'acb' we would shift by 2.
func Decrypt(text string, key string, alphabet string) string {
	k := []rune(key)[0]
	plain := decrypt([]rune(text), k, []rune(alphabet))
	return string(plain)
}

func CrackForce(cipher string, alphabet string, ranker Ranker) []string {
	ps := permutations([]rune(cipher), []rune(alphabet))
	return ranker(toStrings(ps))
}

func CrackFreq(cipher string, alphabet string, frequencies map[string]float64, ranker Ranker) []string {
	mostFreq := sortedFreqs(frequencies)
	top5 := util.Take(mostFreq, 5)
	possibles := make([]string, len(top5))

	for i, v := range top5 {
		possibles[i] = Decrypt(cipher, v, alphabet)
	}

	return ranker(possibles)
}
