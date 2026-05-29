package set1

import "encoding/hex"

// convert hex strings to bytes -> XOR both -> output
func Xor(x, y string) string {
	x_bytes, _ := hex.DecodeString(x)
	y_bytes, _ := hex.DecodeString(y)
	z_bytes := make([]byte, len(x_bytes))
	for i := range z_bytes {
		z_bytes[i] = x_bytes[i] ^ y_bytes[i]
	}
	return hex.EncodeToString(z_bytes)
}
