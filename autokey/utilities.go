package autokey

import "cryptgo/vigenere"

func keyGen(text []rune, primer []rune) []rune {
	key := make([]rune, len(text))
	copy(key, primer)
	copy(key[len(primer):], text)

	return key
}

func decryptChunk(chunk []rune, key string, alphabet string) []rune {
	text := string(chunk)
	return []rune(vigenere.Decrypt(text, key, alphabet))
}

func decrypt(text []rune, primer string, alphabet string) []rune {
	plain := make([]rune, len(text))
	copy(plain, decryptChunk(text[:len(primer)], primer, alphabet))

	key := make([]rune, len(text))
	copy(key, []rune(primer))
	copy(key[len(primer):], plain)

	cursor := len(primer)

	for cursor < len(text) {
		plainchunk := decryptChunk(text[cursor:], string(key[cursor:]), alphabet)

		copy(plain[cursor:], plainchunk)
		copy(key[cursor:], plainchunk)
		cursor += len(plainchunk)
	}

	return plain
}
