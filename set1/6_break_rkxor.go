package set1

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
