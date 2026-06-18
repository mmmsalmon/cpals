package set1

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
	"slices"
	"unicode/utf8"
)

// return the number of differing bits between two byte arrays
func hammingDistance(x, y []byte) int {
	distance := 0
	for i := range x {
		for n := range 8 {
			mask := byte(1 << n)
			if (x[i] & mask) != (y[i] & mask) {
				distance++
			}
		}
	}
	return distance
}

// opens and base64-decodes "txt/6.txt"
// TODO move this into a function that calls everything else
func open6() []byte {
	text, err := os.ReadFile("txt/6.txt")
	if err != nil {
		log.Fatal(err)
	}
	res, _ := base64.StdEncoding.DecodeString(string(text))
	return res
}

// return frequency of each keysize (normalized edit distance)
func findKeysize() map[int]int {
	keySizes := make(map[int]int)
	text := open6()
	for ks := 2; ks < 40; ks++ {
		normalizedDistance := hammingDistance(text[ks:ks*2], text[ks*2:ks*3]) / ks
		keySizes[normalizedDistance]++
	}
	return keySizes
}

// make blocks/chunks containing the nth byte of every chunk
func transpose(ks int) [][]byte {
	chunks := slices.Chunk(open6(), ks)
	transposedBlocks := [][]byte{}
	for range ks {
		for chunk := range chunks {
			transposedBlocks = append(transposedBlocks, chunk)
		}
	}
	return transposedBlocks
}

func SxorBlocks(ks int) {
	var blocks [][]byte = transpose(ks)
	for block := range blocks {
		bloques := make([]byte, ks)
		for k := range 256 {
			key := string([]byte{byte(k)})
			for i := range ks {
				bloques[i] = byte(block) ^ key[0]
			}
			if utf8.ValidString(string(bloques)) {
				fmt.Printf("%s: ", key)
				fmt.Println(string(bloques))
			}
		}
	}
}
