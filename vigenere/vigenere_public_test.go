package vigenere_test

import (
	"cryptgo/caesar"
	"cryptgo/vigenere"
	"testing"
)

const plain = "hello"
const cipher = "hlloo"
const alpha = "helo"

func TestEncrypt(t *testing.T) {
	// Case: one-letter word
	result := vigenere.Encrypt(plain, "e", alpha)
	caesarResult := caesar.Encrypt(plain, "e", alpha)
	if result != caesarResult {
		t.Errorf("Case one-letter word: Expected one letter Vigenere to encrypt to Caesar")
	}

	// Case: Result check
	result = vigenere.Encrypt(plain, "he", alpha)
	if result != cipher {
		t.Errorf("Case result check: Expected %q -- got %q", cipher, result)
	}
}

func TestDecrypt(t *testing.T) {
	// Case: one-letter word
	result := vigenere.Decrypt(cipher, "e", alpha)
	caesarResult := caesar.Decrypt(cipher, "e", alpha)
	if result != caesarResult {
		t.Errorf("Case one-letter word: Expected one letter Vigenere to decrypt to Caesar")
	}

	// Case: Result check
	result = vigenere.Decrypt(cipher, "he", alpha)
	if result != plain {
		t.Errorf("Case result check: Expected %q -- got %q", plain, result)
	}
}
