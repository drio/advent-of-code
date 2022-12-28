package main

import (
	"bufio"
	"fmt"
	"os"
)

const size = 14

func unique(four []string) bool {
	m := map[string]bool{}
	for _, c := range four {
		m[c] = true
	}
	return len(m) == size
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanRunes)
	i := 0
	four := make([]string, size)
	for scanner.Scan() {
		c := scanner.Text()
		if c == "\n" {
			break
		}

		if i < size {
			four[i%size] = c
			i += 1
			continue
		}

		four = append(four[1:], c)

		if unique(four) {
			fmt.Printf("%d\n", i+1)
			break
		}

		i += 1
	}

}
