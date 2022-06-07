package unpad_test

import (
	"cryptgo/internal/ciphers/unpad"
	"cryptgo/internal/util"
	"testing"
)

func TestPKCS7(t *testing.T) {
	table := []struct {
		name     string
		bs       []byte
		expected []byte
	}{
		{name: "Pad empty slice", bs: []byte{0x05, 0x05, 0x05, 0x05, 0x05}, expected: []byte{}},
		{name: "Pad 4 bytes", bs: []byte{0x0, 0x04, 0x04, 0x04, 0x04}, expected: []byte{0x0}},
		{name: "Pad 3 bytes", bs: []byte{0x0, 0x0, 0x03, 0x03, 0x03}, expected: []byte{0x0, 0x0}},
		{name: "Pad 2 bytes", bs: []byte{0x0, 0x0, 0x0, 0x02, 0x02}, expected: []byte{0x0, 0x0, 0x0}},
		{name: "Pad 1 byte", bs: []byte{0x0, 0x0, 0x0, 0x0, 0x01}, expected: []byte{0x0, 0x0, 0x0, 0x0}},
		{name: "Pad multiple of n", bs: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x05, 0x05, 0x05, 0x05, 0x05}, expected: []byte{0x0, 0x0, 0x0, 0x0, 0x0}},
	}

	for _, entry := range table {
		t.Run(entry.name, func(t *testing.T) {
			actual := unpad.PKCS7(entry.bs)

			if !util.AreSlicesEqual(actual, entry.expected) {
				t.Errorf("Actual %v does not equal expected %v", actual, entry.expected)
			}
		})
	}
}

func TestPKCS5(t *testing.T) {
	// Case: Pad in blocks of 8
	actual := unpad.PKCS5([]byte{0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08, 0x08})

	if len(actual) != 0 {
		t.Error("PKCS5 should unpad in increments of 8")
	}
}

func TestOneAndZeroes(t *testing.T) {
	// Case: No 0b1000_0000 found
	actual, err := unpad.OneAndZeroes([]byte{0x01})

	if err == nil {
		t.Errorf("Expected to get an error")
	}

	if actual != nil {
		t.Errorf("Expected to get a nil slice on error cases")
	}

	// Case: Padding
	table := []struct {
		name     string
		bs       []byte
		expected []byte
	}{
		{name: "Empty slice", bs: []byte{0b1000_0000, 0, 0, 0, 0}, expected: []byte{}},
		{name: "1-byte slice", bs: []byte{0xFF, 0b1000_0000, 0, 0, 0}, expected: []byte{0xFF}},
		{name: "2-byte slice", bs: []byte{0xFF, 0xFF, 0b1000_0000, 0, 0}, expected: []byte{0xFF, 0xFF}},
		{name: "3-byte slice", bs: []byte{0xFF, 0xFF, 0xFF, 0b1000_0000, 0}, expected: []byte{0xFF, 0xFF, 0xFF}},
		{name: "4-byte slice", bs: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0b1000_0000}, expected: []byte{0xFF, 0xFF, 0xFF, 0xFF}},
		{name: "5-byte slice", bs: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0b1000_0000, 0, 0, 0, 0}, expected: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
	}

	for _, entry := range table {
		t.Run(entry.name, func(t *testing.T) {
			actual, err := unpad.OneAndZeroes(entry.bs)

			if err != nil {
				t.Errorf("Did not expect to get an error for slice %v", entry.bs)
			}

			if !util.AreSlicesEqual(actual, entry.expected) {
				t.Errorf("Actual %v not equal to expected %v", actual, entry.expected)
			}
		})
	}
}

func TestANSIX923(t *testing.T) {
	table := []struct {
		name     string
		bs       []byte
		n        int
		expected []byte
	}{
		{name: "Empty slice", bs: []byte{0, 0, 0, 0, 0x05}, expected: []byte{}},
		{name: "1-byte slice", bs: []byte{0xFF, 0, 0, 0, 0x04}, expected: []byte{0xFF}},
		{name: "2-byte slice", bs: []byte{0xFF, 0xFF, 0, 0, 0x03}, expected: []byte{0xFF, 0xFF}},
		{name: "3-byte slice", bs: []byte{0xFF, 0xFF, 0xFF, 0, 0x02}, expected: []byte{0xFF, 0xFF, 0xFF}},
		{name: "4-byte slice", bs: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0x01}, expected: []byte{0xFF, 0xFF, 0xFF, 0xFF}},
		{name: "5-byte slice", bs: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0, 0, 0, 0, 0x05}, expected: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
	}

	for _, entry := range table {
		t.Run(entry.name, func(t *testing.T) {
			actual := unpad.ANSIX923(entry.bs)

			if !util.AreSlicesEqual(actual, entry.expected) {
				t.Errorf("Actual %v did not equal %v", actual, entry.expected)
			}
		})
	}
}

func TestW3C(t *testing.T) {
	table := []struct {
		name     string
		bs       []byte
		expected []byte
	}{
		{name: "Empty slice", bs: []byte{0x0, 0x0, 0x0, 0x0, 0x05}, expected: []byte{}},
		{name: "1-byte slice", bs: []byte{0xFF, 0x0, 0x0, 0x0, 0x04}, expected: []byte{0xFF}},
		{name: "2-byte slice", bs: []byte{0xFF, 0xFF, 0x0, 0x0, 0x03}, expected: []byte{0xFF, 0xFF}},
		{name: "3-byte slice", bs: []byte{0xFF, 0xFF, 0xFF, 0x0, 0x02}, expected: []byte{0xFF, 0xFF, 0xFF}},
		{name: "4-byte slice", bs: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0x01}, expected: []byte{0xFF, 0xFF, 0xFF, 0xFF}},
		{name: "5-byte slice", bs: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF, 0x0, 0x0, 0x0, 0x0, 0x05}, expected: []byte{0xFF, 0xFF, 0xFF, 0xFF, 0xFF}},
	}

	for _, entry := range table {
		t.Run(entry.name, func(t *testing.T) {
			actual := unpad.W3C(entry.bs)

			if !util.AreSlicesEqual(actual, entry.expected) {
				t.Errorf("Actual %v did not equal %v", actual, entry.expected)
			}
		})
	}
}
