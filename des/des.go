package des

import (
	"cryptgo/internal/util"
	"cryptgo/padding"
	"errors"
)

type DesOpts = struct {
	Key   []byte
	Pad   padding.Padding
	PByte byte
}

func feistel(bs []byte, ks [][]byte) []byte {
	blocks := util.SplitEvery(bs, 8)
	cipherBlocks := make([][]byte, len(blocks))

	for i, bl := range blocks {
		permuted := initialPermute(toBits(bl))
		split := util.SplitAfterIndices(permuted, 32)

		ln := split[0]
		rn := split[1]

		for n := 0; n < 16; n++ {
			ln, rn = rn, xor(ln, round(rn, ks[n]))
		}

		m := util.Concat(rn, ln)
		cipherBlocks[i] = util.MapSlice(finalPermutation[:], func(i int, _ int) bit {
			return m[i-1]
		})
	}

	return util.Flatten(cipherBlocks)
}

func decrypt(bs []byte, key []byte) []byte {
	subkeys := genSubkeys(toBits(key))
	return feistel(bs, util.Reverse(subkeys))
}

func encrypt(bs []byte, key []byte) []byte {
	subkeys := genSubkeys(toBits(key))
	return feistel(bs, subkeys)
}

// TODO: add modes
func Encrypt(plain []byte, opts DesOpts) ([]byte, error) {
	if len(opts.Key) != 8 {
		return nil, errors.New("Key must be 8 bytes long")
	}

	padded := padders[opts.Pad](plain, opts.PByte)

	if opts.Pad == padding.None && len(padded)%8 != 0 {
		return nil, errors.New("Non-padded messages must have length be a multiple of 8")
	}

	cipher := encrypt(padded, opts.Key)

	return toBytes(cipher), nil
}

func Decrypt(cipher []byte, opts DesOpts) ([]byte, error) {
	if len(opts.Key) != 8 {
		return nil, errors.New("Key must be 8 bytes long")
	}

	unpadded, err := unpadders[opts.Pad](cipher)

	if err != nil {
		return nil, err
	}

	plain := decrypt(unpadded, opts.Key)

	return toBytes(plain), nil
}
