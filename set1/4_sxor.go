package set1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
	"unicode/utf8"
)

// run #3 but loop over a multi-line file
func Sxor(filename string, score []Kv) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("line: %s", line)
		for _, v := range score {
			out := Bxor(line, v.Key)
			if strings.IndexFunc(out, unicode.IsControl) == -1 && utf8.ValidString(out) {
				fmt.Printf("%s: ", v.Key)
				fmt.Println(out)
			}
		}
		fmt.Println()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}
}
