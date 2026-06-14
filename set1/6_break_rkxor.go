package set1

import (
	"log"
	"os"
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

// return frequency of each keysize (normalized edit distance)
func findKeysize(filename string) map[int]int {
	text, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	mDistances := make(map[int]int)
	for ks := 2; ks < 40; ks++ {
		normalizedDistance := hammingDistance(text[ks:ks*2], text[ks*2:ks*3]) / ks
		mDistances[normalizedDistance]++
	}
	return mDistances
}
