package unpad

import (
	"cryptgo/internal/util"
	"errors"
)

func PKCS7(bs []byte) []byte {
	i := util.Last(bs)
	return bs[:len(bs)-int(i)]
}

func PKCS5(bs []byte) []byte {
	return PKCS7(bs)
}

func OneAndZeroes(bs []byte) ([]byte, error) {
	_, i, ok := util.FindLast(bs, func(b byte, _ int) bool {
		return b == 0b1000_0000
	})

	if !ok {
		return nil, errors.New("Could not find initial padding bit")
	}

	return bs[:i], nil
}

func ANSIX923(bs []byte) []byte {
	i := util.Last(bs)
	return bs[:len(bs)-int(i)]
}

func W3C(bs []byte) []byte {
	i := util.Last(bs)
	return bs[:len(bs)-int(i)]
}
