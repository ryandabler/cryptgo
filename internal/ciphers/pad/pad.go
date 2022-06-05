package pad

import (
	"cryptgo/internal/util"
	"errors"
	"fmt"
	"math"
)

const MAX_PKCS5_BLOCK_SIZE = 8

// PKCS7 pad a slice of bytes to the next multiple of `n`.
//
// The pad value is the number of bytes being padded, and
// so the maximum value `n` can be is 0xFF (255). If over,
// an error will be returned. If the byte slice is already
// a multiple of the block, an entire block will be added.
func PKCS7(bs []byte, n int) ([]byte, error) {
	if n > math.MaxUint8 {
		return nil, errors.New(fmt.Sprintf("Block size must be less than or equal to %d: got %d", math.MaxUint8, n))
	}

	pad := n - (len(bs) % n)

	padded := util.Concat(
		bs,
		util.Fill(make([]byte, pad), byte(pad)),
	)

	return padded, nil
}

// PKCS5 pad a slice of bytes to the next multiple of 8.
//
// Uses PKCS7 logic under the hood, but fixes the block
// size to be 8 bytes.
func PKCS5(bs []byte) ([]byte, error) {
	return PKCS7(bs, MAX_PKCS5_BLOCK_SIZE)
}

// OneAndZeroes pad a slice of bytes to the next multiple of `n`.
//
// Will pad the given slice with a byte containing one `1` at
// the front and zeroes for every other padded bit. If the slice
// is a multiple of `n` already, an entire block of padding will
// be added.
func OneAndZeroes(bs []byte, n int) []byte {
	pad := n - (len(bs) % n)

	padding := util.Fill(make([]byte, pad), 0)
	padding[0] = 0b1000_0000

	return util.Concat(bs, padding)
}

// ANSIX923 pad a slice of bytes to the next multiple of `n`.
//
// Padding will be all zeroes with the last byte of the padding
// representing how many padded bytes were added. Because of
// this, if `n` is greater than 255 an error will be returned.
func ANSIX923(bs []byte, n int) ([]byte, error) {
	if n > math.MaxUint8 {
		return nil, errors.New(fmt.Sprintf("Block size must be less than or equal to %d: got %d", math.MaxUint8, n))
	}

	pad := n - (len(bs) % n)

	padding := util.Fill(make([]byte, pad), 0)
	padding[len(padding)-1] = byte(pad)

	return util.Concat(bs, padding), nil
}

// W3C pad a slice of bytes to the next multiple of `n`
// using specified by `p`.
//
// Padding bytes will be specified by the caller, but the
// last padded byte will represent the number of bytes
// padded. As such, the block size cannot be greater than
// 255 or an error will be returned.
func W3C(bs []byte, n int, p byte) ([]byte, error) {
	if n > math.MaxUint8 {
		return nil, errors.New(fmt.Sprintf("Block size must be less than or equal to %d: got %d", math.MaxUint8, n))
	}

	pad := n - (len(bs) % n)

	padding := util.Fill(make([]byte, pad), p)
	padding[len(padding)-1] = byte(pad)

	return util.Concat(bs, padding), nil
}
