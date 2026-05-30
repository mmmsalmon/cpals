package set1

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
)

func RepeatingXOR(filename, key string) {
	text, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	var xorLine []byte
	for i, c := range text {
		xorLine = append(xorLine, byte(c)^key[i%len(key)])
	}
	fmt.Println(hex.EncodeToString(xorLine))
}
