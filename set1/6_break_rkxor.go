package set1

import (
	"encoding/base64"
	"log"
	"os"
	"slices"
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
func FindKeysize() map[int]int {
	keySizes := make(map[int]int)
	text := open6()
	for ks := 2; ks < 40; ks++ {
		normalizedDistance := hammingDistance(text[ks:ks*2], text[ks*2:ks*3]) / ks
		keySizes[normalizedDistance]++
	}
	return keySizes
}

func transpose(ks int) {
	chunks := slices.Chunk(open6(), ks)
	transposedBlocks := []byte{}
	for key := range ks {
		for chunk := range chunks {
			transposedBlocks[key] += chunk[key]
		}
	}
}
