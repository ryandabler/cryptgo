package des_test

import (
	"cryptgo/des"
	"cryptgo/internal/util"
	"testing"
)

var msg []byte = []byte{0b0000_0001, 0b0010_0011, 0b0100_0101, 0b0110_0111, 0b1000_1001, 0b1010_1011, 0b1100_1101, 0b1110_1111}
var key []byte = []byte{0b0001_0011, 0b0011_0100, 0b0101_0111, 0b0111_1001, 0b1001_1011, 0b1011_1100, 0b1101_1111, 0b1111_0001}
var cipher []byte = []byte{0b1000_0101, 0b1110_1000, 0b0001_0011, 0b0101_0100, 0b0000_1111, 0b0000_1010, 0b1011_0100, 0b0000_0101}

func TestEncrypt(t *testing.T) {
	// Case: encrypting message
	actual, _ := des.Encrypt(msg, key)

	if !util.AreSlicesEqual(cipher, actual) {
		t.Errorf("Cipher text does not match expected value")
	}

	// Case: Key is not of right length
	actual, err := des.Encrypt(msg, key[:len(key)-1])

	if err == nil {
		t.Errorf("Expected error returned when key is not 8 bytes long")
	}

	if actual != nil {
		t.Errorf("Ciphertext should be nil when error is returned")
	}
}

func TestDecrypt(t *testing.T) {
	// Case: decrypting message
	actual, _ := des.Decrypt(cipher, key)

	if !util.AreSlicesEqual(msg, actual) {
		t.Errorf("Plain text does not match expected value")
	}

	// Case: Key is not of right length
	plain, err := des.Decrypt(cipher, key[:len(key)-1])

	if err == nil {
		t.Errorf("Expected error returned when key is not 8 bytes long")
	}

	if plain != nil {
		t.Errorf("Plaintext should be nil when error is returned")
	}
}
