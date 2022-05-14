package des

import (
	"cryptgo/internal/util"
)

// Pad slice of bytes `bs` to be an even multiple of int `mod`.
// Any padded values should be of value byte `p`.
func pad(bs []byte, mod int, p byte) []byte {
	for len(bs)%mod != 0 {
		bs = append(bs, p)
	}

	return bs
}

func initialPermute(bs []bit) []bit {
	return util.MapSlice(initialPermutation[:], func(i int, _ int) bit {
		return bs[i-1]
	})
}

// Retrieve row of SBox designated by bit slice `bs`
func sBoxRow(bs []bit) byte {
	first := bs[0]
	last := bs[5]

	return (first << 1) | last
}

// Retrieve column of SBox designated by bit slice `bs`
func sBoxCol(bs []bit) byte {
	var b byte
	bs = bs[1:5]

	for i, v := range bs {
		b |= (v << (len(bs) - 1 - i))
	}

	return b
}

// Retrieve 4-bit slice designated by bit slice `bs`.
// Due to SBox values being stored linearly, we must
// map each (row, col) pair to an index in a slice.
func sBox(bs []bit, n int) []bit {
	col := sBoxCol(bs)
	row := sBoxRow(bs)
	val := sboxes[n][row*16+col]

	return toTruncatedBits(val, 4)
}

// Perform the round function on bit slice `bs` with key `k`.
// `bs`` should be 4 bytes, `k`` should be 6 bytes
func round(bs []bit, k []bit) []bit {
	stage1 := xor(expand(bs), k)
	stage2 := util.SplitEvery(stage1, 6)
	stage3 := make([]bit, len(bs))

	for i, v := range stage2 {
		copy(stage3[i*4:(i+1)*4], sBox(v, i))
	}

	return util.MapSlice(permutation[:], func(i int, _ int) bit {
		return stage3[i-1]
	})
}

// Index-wise XOR operation on two equal-length bit-slices
func xor(bs []bit, cs []bit) []bit {
	return util.MapSlice(bs, func(b bit, i int) bit {
		return b ^ cs[i]
	})
}

// Expand a 32-bit slice `bs` to 48 bits
func expand(bs []bit) []bit {
	return util.MapSlice(eTable[:], func(i int, _ int) bit {
		return bs[i-1]
	})
}
