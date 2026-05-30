package set1

import (
	"encoding/hex"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
	"unicode"
)

// for sorting a map later
type kv struct {
	Key   string
	Value int
}

// take a sample text and 'score' it by measuring character frequency
// 'filename' should point to a file with a single line of text
func scoring(filename string) []kv {
	text, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	convtext := strings.ToUpper(string(text))
	text = []byte(convtext)

	m := make(map[string]int)
	for _, val := range text {
		m[string(val)]++
	}

	var ss []kv
	for k, v := range m {
		ss = append(ss, kv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool { return ss[i].Value > ss[j].Value })

	return ss
}

// decode XOR-ciphered message
func Bxor(x, y string) string {
	x_bytes, _ := hex.DecodeString(x)
	key := []byte(y)
	z_bytes := make([]byte, len(x_bytes))
	for i := range z_bytes {
		z_bytes[i] = x_bytes[i] ^ key[0]
	}
	return string(z_bytes)
}

func ScoreDecipher(filename, hexstring string) {
	score := scoring(filename)
	for _, v := range score {
		out := Bxor(hexstring, v.Key)
		if strings.IndexFunc(out, unicode.IsControl) == -1 {
			fmt.Printf("%s: ", v.Key)
			fmt.Println(out)
		}
	}
}
