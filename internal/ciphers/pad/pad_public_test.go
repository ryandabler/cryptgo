package pad_test

import (
	"cryptgo/internal/ciphers/pad"
	"cryptgo/internal/util"
	"math"
	"testing"
)

func TestPKCS7(t *testing.T) {
	// Case: Block size larger than maximum uint8 value
	res, err := pad.PKCS7([]byte{}, math.MaxUint8+1)

	if res != nil {
		t.Errorf("Expected padded slice to be nil")
	}

	if err == nil {
		t.Errorf("Expected error expressing inability to pad")
	}

	// Case: Padding
	n := 5
	table := []struct {
		name     string
		bs       []byte
		n        int
		expected []byte
	}{
		{name: "Pad empty slice", bs: []byte{}, n: n, expected: []byte{0x05, 0x05, 0x05, 0x05, 0x05}},
		{name: "Pad 4 bytes", bs: []byte{0x0}, n: n, expected: []byte{0x0, 0x04, 0x04, 0x04, 0x04}},
		{name: "Pad 3 bytes", bs: []byte{0x0, 0x0}, n: n, expected: []byte{0x0, 0x0, 0x03, 0x03, 0x03}},
		{name: "Pad 2 bytes", bs: []byte{0x0, 0x0, 0x0}, n: n, expected: []byte{0x0, 0x0, 0x0, 0x02, 0x02}},
		{name: "Pad 1 byte", bs: []byte{0x0, 0x0, 0x0, 0x0}, n: n, expected: []byte{0x0, 0x0, 0x0, 0x0, 0x01}},
		{name: "Pad multiple of n", bs: []byte{0x0, 0x0, 0x0, 0x0, 0x0}, n: n, expected: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x05, 0x05, 0x05, 0x05, 0x05}},
	}

	for _, entry := range table {
		t.Run(entry.name, func(t *testing.T) {
			actual, err := pad.PKCS7(entry.bs, entry.n)

			if err != nil {
				t.Errorf("Did not expect err for test %q", entry.name)
			}

			if !util.AreSlicesEqual(actual, entry.expected) {
				t.Errorf("Actual %v does not equal expected %v", actual, entry.expected)
			}
		})
	}
}

func TestPKCS5(t *testing.T) {
	// Case: Pad in blocks of 8
	actual, err := pad.PKCS5([]byte{})

	if err != nil {
		t.Error("PKCS5 should not have errors")
	}

	if len(actual) != 8 {
		t.Error("PKCS5 should pad in increments of 8")
	}

	if !util.Every(actual, func(b byte, _ int) bool {
		return b == 0x08
	}) {
		t.Error("PKCS5 should pad using 0x08 for an empty slice")
	}
}

func TestOneAndZeroes(t *testing.T) {
	// Case: Padding
	n := 5
	table := []struct {
		name     string
		bs       []byte
		n        int
		expected []byte
	}{
		{name: "Empty slice", bs: []byte{}, n: n, expected: []byte{0b1000_0000, 0, 0, 0, 0}},
		{name: "1-byte slice", bs: []byte{0xFF}, n: n, expected: []byte{0xFF, 0b1000_0000, 0, 0, 0}},
		{name: "2-byte slice", bs: []byte{0xFF, 0xFF}, n: n, expected: []byte{0xFF, 0xFF, 0b1000_0000, 0, 0}},
		{name: "3-byte slice", bs: []byte{0xFF, 0xFF, 0xFF}, n: n, expected: []byte{0xFF, 0xFF, 0xFF, 0b1000_0000, 0}},
		{name: "4-byte slice", bs: []byte{0xFF, 0xFF, 0xFF, 0xFF}, n: n, expected: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0b1000_0000}},
		{name: "5-byte slice", bs: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, n: n, expected: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0b1000_0000, 0, 0, 0, 0}},
	}

	for _, entry := range table {
		t.Run(entry.name, func(t *testing.T) {
			actual := pad.OneAndZeroes(entry.bs, entry.n)

			if !util.AreSlicesEqual(actual, entry.expected) {
				t.Errorf("Actual %v not equal to expected %v", actual, entry.expected)
			}
		})
	}
}

func TestANSIX923(t *testing.T) {
	// Case: Block size larger than maximum uint8 value
	res, err := pad.ANSIX923([]byte{}, math.MaxUint8+1)

	if res != nil {
		t.Errorf("Expected padded slice to be nil")
	}

	if err == nil {
		t.Errorf("Expected error expressing inability to pad")
	}

	// Case: Padding
	n := 5
	table := []struct {
		name     string
		bs       []byte
		n        int
		expected []byte
	}{
		{name: "Empty slice", bs: []byte{}, n: n, expected: []byte{0, 0, 0, 0, 0x05}},
		{name: "1-byte slice", bs: []byte{0xFF}, n: n, expected: []byte{0xFF, 0, 0, 0, 0x04}},
		{name: "2-byte slice", bs: []byte{0xFF, 0xFF}, n: n, expected: []byte{0xFF, 0xFF, 0, 0, 0x03}},
		{name: "3-byte slice", bs: []byte{0xFF, 0xFF, 0xFF}, n: n, expected: []byte{0xFF, 0xFF, 0xFF, 0, 0x02}},
		{name: "4-byte slice", bs: []byte{0xFF, 0xFF, 0xFF, 0xFF}, n: n, expected: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0x01}},
		{name: "5-byte slice", bs: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, n: n, expected: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0, 0, 0, 0, 0x05}},
	}

	for _, entry := range table {
		t.Run(entry.name, func(t *testing.T) {
			actual, err := pad.ANSIX923(entry.bs, entry.n)

			if err != nil {
				t.Errorf("Did not expect error for test %q", entry.name)
			}

			if !util.AreSlicesEqual(actual, entry.expected) {
				t.Errorf("Actual %v did not equal %v", actual, entry.expected)
			}
		})
	}
}

func TestW3C(t *testing.T) {
	// Case: Block size larger than maximum uint8 value
	res, err := pad.W3C([]byte{}, math.MaxUint8+1, 0)

	if res != nil {
		t.Errorf("Expected padded slice to be nil")
	}

	if err == nil {
		t.Errorf("Expected error expressing inability to pad")
	}

	// Case: Padding
	n := 5
	p := byte(0x13)
	table := []struct {
		name     string
		bs       []byte
		n        int
		p        byte
		expected []byte
	}{
		{name: "Empty slice", bs: []byte{}, n: n, p: p, expected: []byte{p, p, p, p, 0x05}},
		{name: "1-byte slice", bs: []byte{0xFF}, n: n, p: p, expected: []byte{0xFF, p, p, p, 0x04}},
		{name: "2-byte slice", bs: []byte{0xFF, 0xFF}, n: n, p: p, expected: []byte{0xFF, 0xFF, p, p, 0x03}},
		{name: "3-byte slice", bs: []byte{0xFF, 0xFF, 0xFF}, n: n, p: p, expected: []byte{0xFF, 0xFF, 0xFF, p, 0x02}},
		{name: "4-byte slice", bs: []byte{0xFF, 0xFF, 0xFF, 0xFF}, n: n, p: p, expected: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0x01}},
		{name: "5-byte slice", bs: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF}, n: n, p: p, expected: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, p, p, p, p, 0x05}},
	}

	for _, entry := range table {
		t.Run(entry.name, func(t *testing.T) {
			actual, err := pad.W3C(entry.bs, entry.n, entry.p)

			if err != nil {
				t.Errorf("Did not expect an error for test %q", entry.name)
			}

			if !util.AreSlicesEqual(actual, entry.expected) {
				t.Errorf("Actual %v did not equal %v", actual, entry.expected)
			}
		})
	}
}
