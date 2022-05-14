package des

import (
	"cryptgo/internal/util"
	"errors"
)

func encrypt(bs []byte, key []byte) []byte {
	subkeys := genSubkeys(toBits(key))
	blocks := util.SplitEvery(bs, 8)
	cipherBlocks := make([][]byte, len(blocks))

	for i, bl := range blocks {
		permuted := initialPermute(toBits(bl))
		split := util.SplitAfterIndices(permuted, 32)

		ln := split[0]
		rn := split[1]

		for n := 0; n < 16; n++ {
			ln, rn = rn, xor(ln, round(rn, subkeys[n]))
		}

		m := util.Concat(rn, ln)
		cipherBlocks[i] = util.MapSlice(finalPermutation[:], func(i int, _ int) bit {
			return m[i-1]
		})
	}

	return util.Flatten(cipherBlocks)
}

// TODO: add padding styles, add modes
func Encrypt(plain string, key string) (string, error) {
	bKey := []byte(key)

	if len(bKey) != 8 {
		return "", errors.New("Key must be 16 bytes long")
	}

	bPlain := []byte(plain)
	bPlain = pad(bPlain, 8, 0)
	cipher := encrypt(bPlain, bKey)

	return string(cipher), nil
}
