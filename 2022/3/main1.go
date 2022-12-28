package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	p := genPriorities()
	fmt.Printf("%d %d %d %d\n", p["a"], p["z"], p["A"], p["Z"])
	sum := 0

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		inFirst := map[string]bool{}
		chars := strings.Split(line, "")

		firstHalf := chars[0 : len(line)/2]
		for _, c := range firstHalf {
			inFirst[c] = true
		}

		secondHalf := chars[len(line)/2:]
		for idx, c := range secondHalf {
			if _, ok := inFirst[c]; ok {
				sum += p[c]
				fmt.Printf("*%s*(%d)", secondHalf[idx], p[c])
				break
			} else {
				fmt.Printf("%s", secondHalf[idx])
			}
		}
		fmt.Printf("\n")
	}

	fmt.Printf("%d", sum)
}

func genPriorities() map[string]int {
	priorities := map[string]int{}
	i := 1
	for r := 'a'; r <= 'z'; r++ {
		priorities[string(r)] = i
		i += 1
	}
	for r := 'A'; r <= 'Z'; r++ {
		priorities[string(r)] = i
		i += 1
	}
	return priorities
}
