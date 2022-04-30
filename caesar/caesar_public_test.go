package caesar_test

import (
	"cryptgo/caesar"
	"testing"
)

const plain = "hello"
const cipher = "elooh"
const alpha = "helo"

func identity[T any](ss []T) []T {
	return ss
}

// Package level variable to force no compiler optimizations when
// running benchmarking code
var results []string

func TestEncrypt(t *testing.T) {
	// Case: shifting by first letter of alphabet should be
	// a no-op
	result := caesar.Encrypt(plain, "h", alpha)
	if result != plain {
		t.Errorf("Case no-op: Expected %q -- got %q", plain, result)
	}

	// Case: shifting should wrap left to right
	result = caesar.Encrypt(plain, "e", alpha)
	if result != cipher {
		t.Errorf("Case wrap: Expected %q -- got %q", cipher, result)
	}
}

func TestDecrypt(t *testing.T) {
	// Case: shifting by first letter of alphabet should be
	// a no-op
	result := caesar.Decrypt(cipher, "h", alpha)
	if result != cipher {
		t.Errorf("Case no-op: Expected %q -- got %q", cipher, result)
	}

	// Case: shifting should wrap right to left
	result = caesar.Decrypt(cipher, "e", alpha)
	if result != plain {
		t.Errorf("Case wrap: Expected %q -- got %q", plain, result)
	}
}

func TestCrackForce(t *testing.T) {
	alphaLen := len([]rune(alpha))

	// Case: should return all permutations based on length of alphabet
	results := caesar.CrackForce(cipher, alpha, identity[string])
	if len(results) != alphaLen {
		t.Errorf("Case all permutations: expected %d results -- got %d", alphaLen, len(results))
	}

	// Case: `ranker` function should be called
	calls := 0
	caesar.CrackForce(cipher, alpha, func(ss []string) []string {
		calls++
		return ss
	})
	if calls == 0 {
		t.Errorf("Case ranker called: Expected `ranker` function to be called")
	}
}

func TestCrackFreq(t *testing.T) {
	// Case: `ranker` function should be called
	calls := 0
	caesar.CrackFreq(cipher, alpha, func(ss []string) []string {
		calls++
		return ss
	})
	if calls == 0 {
		t.Errorf("Case ranker called: Expected `ranker` function to be called")
	}

	// Case: `ranker` function should be called with top potential results
	var args []string
	caesar.CrackFreq(cipher, alpha, func(ss []string) []string {
		args = ss
		return ss
	})
	if args == nil {
		t.Errorf("Case ranker args: Expected `ranker` function to be called with type `[]string`")
	}
}

func BenchmarkCrackForce(b *testing.B) {
	alpha := "abcdefghijklmnopqrstuvwxyz "
	cipher := "jlyhcphcdceuhdncjlyhcphcdceuhdn"

	var r []string

	for i := 0; i < b.N; i++ {
		r = caesar.CrackForce(cipher, alpha, func(ss []string) []string { return ss })
	}

	results = r
}
