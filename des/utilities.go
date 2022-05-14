package des

type bit = byte
type endian int

const (
	bigEnd endian = iota
	litEnd
)

// Get a bit in byte `b` at position `p` based on endianness
// If big endian, that means we want to get the position starting
// from the LEFT; little endian means we want bit from the RIGHT
func getBitAtPos(b byte, p int, endian endian) bit {
	shift := p

	if endian == bigEnd {
		// Indexing starts at 0, so we use 7 instead of 8
		shift = 7 - p
	}

	return (b >> shift) & 1
}

// Convert an slice of bytes to a slice of bits
func toBits(bs []byte) []bit {
	bits := make([]bit, len(bs)*8)

	for i, v := range bs {
		for j := 0; j < 8; j++ {
			bits[i*8+j] = getBitAtPos(v, j, bigEnd)
		}
	}

	return bits
}

// Convert an int `n` to a slice of bits with a required minimum
// padding (padded 0) of `p`
func toTruncatedBits(n int, p int) []bit {
	bs := make([]bit, p)

	for i := p; i > 0; i-- {
		bs[i-1] = bit(n & 1)
		n >>= 1
	}

	return bs
}