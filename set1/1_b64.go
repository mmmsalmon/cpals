package set1

import (
	"encoding/base64"
	"encoding/hex"
)

// convert hex string -> bytes -> base64 string
func B64(x string) string {
	hexb, _ := hex.DecodeString(x)
	return base64.StdEncoding.EncodeToString(hexb)
}
