package des_test

import (
	"cryptgo/des"
	"cryptgo/internal/util"
	"cryptgo/padding"
	"testing"
)

var msg []byte = []byte{0b0000_0001, 0b0010_0011, 0b0100_0101, 0b0110_0111, 0b1000_1001, 0b1010_1011, 0b1100_1101, 0b1110_1111}
var key []byte = []byte{0b0001_0011, 0b0011_0100, 0b0101_0111, 0b0111_1001, 0b1001_1011, 0b1011_1100, 0b1101_1111, 0b1111_0001}
var cipher []byte = []byte{0b1000_0101, 0b1110_1000, 0b0001_0011, 0b0101_0100, 0b0000_1111, 0b0000_1010, 0b1011_0100, 0b0000_0101}

func defaultOpts() des.DesOpts {
	return des.DesOpts{
		Key:   key,
		Pad:   padding.None,
		PByte: 0,
	}
}

func TestEncrypt(t *testing.T) {
	opts := defaultOpts()

	// Case: encrypting message
	actual, _ := des.Encrypt(msg, opts)

	if !util.AreSlicesEqual(cipher, actual) {
		t.Errorf("Cipher text does not match expected value")
	}

	// Case: Key is not of right length
	opts = defaultOpts()
	opts.Key = opts.Key[:len(key)-1]

	actual, err := des.Encrypt(msg, opts)

	if err == nil {
		t.Errorf("Expected error returned when key is not 8 bytes long")
	}

	if actual != nil {
		t.Errorf("Ciphertext should be nil when error is returned")
	}

	// Case: No padding of message whose length isn't a multiple of 8
	opts = defaultOpts()
	opts.Pad = padding.None

	actual, err = des.Encrypt(msg[:len(msg)-1], opts)

	if err == nil {
		t.Errorf("Expected error returned when message length is not multiple of 8 and no padding is performed")
	}

	if actual != nil {
		t.Error("Ciphertext should be nil when error is returned")
	}
}

func TestDecrypt(t *testing.T) {
	opts := defaultOpts()

	// Case: decrypting message
	actual, _ := des.Decrypt(cipher, opts)

	if !util.AreSlicesEqual(msg, actual) {
		t.Errorf("Plain text does not match expected value")
	}

	// Case: Key is not of right length
	opts = defaultOpts()
	opts.Key = opts.Key[:len(key)-1]

	plain, err := des.Decrypt(cipher, opts)

	if err == nil {
		t.Errorf("Expected error returned when key is not 8 bytes long")
	}

	if plain != nil {
		t.Errorf("Plaintext should be nil when error is returned")
	}
}
