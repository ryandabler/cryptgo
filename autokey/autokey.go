package autokey

import (
	"cryptgo/vigenere"
)

func Encrypt(text string, primer string, alphabet string) string {
	key := keyGen([]rune(text), []rune(primer))
	return vigenere.Encrypt(text, string(key), alphabet)
}

func Decrypt(text string, primer string, alphabet string) string {
	plain := decrypt([]rune(text), primer, alphabet)
	return string(plain)
}
