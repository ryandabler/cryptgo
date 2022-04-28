package onetimepad

import (
	"cryptgo/internal/util"
	"cryptgo/vigenere"
	"crypto/rand"
	"errors"
	"math/big"
)

func randInts(n int, upper int) ([]int, error) {
	ints := make([]int, n)

	for i := 0; i < n; i++ {
		v, err := rand.Int(rand.Reader, big.NewInt(int64(upper)))

		if err != nil {
			return nil, err
		}

		ints[i] = int(v.Int64())
	}

	return ints, nil
}

func Encrypt(text string, key string, alphabet string) (string, string, error) {
	runes := []rune(text)
	alphaR := []rune(alphabet)
	ekey := []rune(key)

	// Non-empty key supplied so check that it is of valid length (must be at least as long
	// as plaintext)
	if 0 < len(ekey) && len(ekey) < len(runes) {
		return "", "", errors.New("Invalid encryption key supplied: less than length of plaintext")
	}

	// Empty key supplied, so we must generate the key to use for encryption
	if len(ekey) == 0 {
		ints, err := randInts(len(runes), len(alphaR))

		if err != nil {
			return "", "", errors.New("Could not generate encryption key")
		}

		ekey = util.MapSlice(ints, func(i int) rune { return alphaR[i] })
	}

	key = string(ekey)
	cipher := vigenere.Encrypt(text, key, alphabet)
	return string(cipher), key, nil
}

func Decrypt(text string, key string, alphabet string) string {
	return vigenere.Decrypt(text, key, alphabet)
}
