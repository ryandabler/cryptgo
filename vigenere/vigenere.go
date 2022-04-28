package vigenere

import "cryptgo/internal/util"

func encrypt(rs []rune, ks []rune, alphabet []rune) []rune {
	cipher := make([]rune, len(rs))
	mod := len(alphabet)

	for i, r := range rs {
		k := ks[i%len(ks)]
		shift := util.IndexOf(ks, k)
		pos := util.IndexOf(alphabet, r)
		newPos := (pos - shift + mod) % mod

		cipher[i] = alphabet[newPos]
	}

	return cipher
}

func decrypt(rs []rune, ks []rune, alphabet []rune) []rune {
	plain := make([]rune, len(rs))
	mod := len(alphabet)

	for i, r := range rs {
		k := ks[i%len(ks)]
		shift := util.IndexOf(ks, k)
		pos := util.IndexOf(alphabet, r)
		newPos := (pos + shift) % mod

		plain[i] = alphabet[newPos]
	}

	return plain
}

// Encrypt a string `text` of alphabet `alphabet` using
// keyword `key`.
//
// Each character of `text` will be left-shifted by the
// numerical value of a letter in `key` as it exists in
// the alphabet. This means the magnitude of left-shifting
// is based on where `k` appears in the alphabet. Thus if
// we shift by 'b' in alphabet 'abc' we would shift by 1
// whereas in the alphabet 'acb' we would shift by 2.
func Encrypt(text string, key string, alphabet string) string {
	k := []rune(key)
	cipher := encrypt([]rune(text), k, []rune(alphabet))
	return string(cipher)
}

// Decrypt a string `text` of alphabet `alphabet` shifting right
// by values of letters in keyword `key`.
//
// Each character of `text` will be right-shifted by the
// numerical value of a letter in `key` as it exists in the
// alphabet. This means the magnitude of right-shifting is
// based on where `k` appears in the alphabet. Thus if we
// shift by 'b' in alphabet 'abc' we would shift by 1 whereas
// in the alphabet 'acb' we would shift by 2.
func Decrypt(text string, key string, alphabet string) string {
	plain := decrypt([]rune(text), []rune(key), []rune(alphabet))
	return string(plain)
}
