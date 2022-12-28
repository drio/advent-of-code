package main

import (
	"bufio"
	"fmt"
	"os"
)

type stack struct {
	crates []rune
}

func (s *stack) addToBottom(r rune) {
	s.crates = append([]rune{r}, s.crates...)
}

func (s *stack) push(r []rune) {
	s.crates = append(s.crates, r...)
}

func (s *stack) pop(n int) []rune {
	// TODO: this assumes we always have something in the stack
	r := s.crates[len(s.crates)-n : len(s.crates)]
	s.crates = s.crates[0 : len(s.crates)-n]
	return r
}

func main() {
	stacks := make([]stack, 9)
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
	for string(scanner.Text()[1]) != "1" {
		for i, r := range scanner.Text() {
			if r != ' ' && r != '[' && r != ']' {
				stacks[i/4].addToBottom(r)
			}
		}
		scanner.Scan()
	}

	// empty line
	scanner.Scan()

	for scanner.Scan() {
		var num, from, to int
		fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &num, &from, &to)
		stacks[to-1].push(stacks[from-1].pop(num))
	}

	for _, s := range stacks {
		fmt.Printf("%s", string(s.pop(1)))
	}
}
