package set1

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode/utf8"
)

// run #3 but loop over a multi-line file
func Sxor(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Printf("line: %s", line)
		for k := range 256 {
			key := string([]byte{byte(k)})
			out := Bxor(line, key)
			if utf8.ValidString(out) {
				fmt.Printf("%s: ", key)
				fmt.Println(out)
			}
		}
		fmt.Println()
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
	}
}
