package des

import (
	"cryptgo/internal/util"
)

const NUM_SUBKEYS = 16

// Shift elements in slice `b` to the left by `s` places.
// Whatever is removed during the shift will be added on
// to the right of the slice to fill in the "gaps" from the
// shift
func leftShiftWrap[T any](b []T, s int) []T {
	newB := make([]T, len(b))

	copy(newB, b[s:])
	copy(newB[len(b)-s:], b[:s])

	return newB
}

// Generate a slice of subkeys based off `key`.
// `key` should be 64 bits and should return 16x 48 bit keys
func genSubkeys(key []bit) [][]bit {
	subkeys := make([][]bit, NUM_SUBKEYS)
	subkey := util.MapSlice(permutedChoice1[:], func(i int, _ int) bit {
		return key[i-1]
	})

	// Split permuted subkey into two halves
	splitSubkey := util.SplitAfterIndices(subkey, len(subkey)/2)

	for i := 0; i < len(subkeys); i++ {
		shift := keyShifts[i]
		c := leftShiftWrap(splitSubkey[0], shift)
		d := leftShiftWrap(splitSubkey[1], shift)
		prekey := util.Concat(c, d)
		subkeys[i] = util.MapSlice(permutedChoice2[:], func(i int, _ int) bit {
			return prekey[i-1]
		})

		splitSubkey[0] = c
		splitSubkey[1] = d
	}

	return subkeys
}
