package pad

import (
	"cryptgo/internal/util"
	"errors"
	"fmt"
	"math"
)

const MAX_PKCS5_BLOCK_SIZE = 8

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

func PKCS5(bs []byte) ([]byte, error) {
	return PKCS7(bs, MAX_PKCS5_BLOCK_SIZE)
}

func OneAndZeroes(bs []byte, n int) []byte {
	pad := n - (len(bs) % n)

	padding := util.Fill(make([]byte, pad), 0)
	padding[0] = 0b1000_0000

	return util.Concat(bs, padding)
}

func ANSIX923(bs []byte, n int) ([]byte, error) {
	if n > math.MaxUint8 {
		return nil, errors.New(fmt.Sprintf("Block size must be less than or equal to %d: got %d", math.MaxUint8, n))
	}

	pad := n - (len(bs) % n)

	padding := util.Fill(make([]byte, pad), 0)
	padding[len(padding)-1] = byte(pad)

	return util.Concat(bs, padding), nil
}

func W3C(bs []byte, n int, p byte) ([]byte, error) {
	if n > math.MaxUint8 {
		return nil, errors.New(fmt.Sprintf("Block size must be less than or equal to %d: got %d", math.MaxUint8, n))
	}

	pad := n - (len(bs) % n)

	padding := util.Fill(make([]byte, pad), p)
	padding[len(padding)-1] = byte(pad)

	return util.Concat(bs, padding), nil
}
