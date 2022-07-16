package padding

type Padding int

const (
	PKCS7 Padding = iota
	PKCS5
	OneAndZeroes
	ANSIX923
	W3C
	None
)
