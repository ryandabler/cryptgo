package des_test

import (
	"cryptgo/des"
	"testing"
)

func TestEncrypt(t *testing.T) {
	// Case: encrypting message
	msg := string([]byte{0b0000_0001, 0b0010_0011, 0b0100_0101, 0b0110_0111, 0b1000_1001, 0b1010_1011, 0b1100_1101, 0b1110_1111})
	k := string([]byte{0b0001_0011, 0b0011_0100, 0b0101_0111, 0b0111_1001, 0b1001_1011, 0b1011_1100, 0b1101_1111, 0b1111_0001})

	cipher, _ := des.Encrypt(msg, k)

	if cipher != string([]byte{0b1000_0101, 0b1110_1000, 0b0001_0011, 0b0101_0100, 0b0000_1111, 0b0000_1010, 0b1011_0100, 0b0000_0101}) {
		t.Errorf("Cipher text does not match expected value")
	}

	// Case: Key is not of right length
	cipher, err := des.Encrypt(msg, string([]byte{0b0001_0011, 0b0011_0100, 0b0101_0111, 0b0111_1001, 0b1001_1011, 0b1011_1100, 0b1101_1111}))

	if err == nil {
		t.Errorf("Expected error returned when key is not 8 bytes long")
	}

	if cipher != "" {
		t.Errorf("Ciphertext should be empty string when error is returned")
	}
}

func TestDecrypt(t *testing.T) {
	// Case: decrypting message
	cipher := string([]byte{0b1000_0101, 0b1110_1000, 0b0001_0011, 0b0101_0100, 0b0000_1111, 0b0000_1010, 0b1011_0100, 0b0000_0101})
	k := string([]byte{0b0001_0011, 0b0011_0100, 0b0101_0111, 0b0111_1001, 0b1001_1011, 0b1011_1100, 0b1101_1111, 0b1111_0001})

	plain, _ := des.Decrypt(cipher, k)

	if plain != string([]byte{0b0000_0001, 0b0010_0011, 0b0100_0101, 0b0110_0111, 0b1000_1001, 0b1010_1011, 0b1100_1101, 0b1110_1111}) {
		t.Errorf("Plain text does not match expected value")
	}

	// Case: Key is not of right length
	plain, err := des.Decrypt(cipher, string([]byte{0b0001_0011, 0b0011_0100, 0b0101_0111, 0b0111_1001, 0b1001_1011, 0b1011_1100, 0b1101_1111}))

	if err == nil {
		t.Errorf("Expected error returned when key is not 8 bytes long")
	}

	if plain != "" {
		t.Errorf("Plaintext should be empty string when error is returned")
	}
}
