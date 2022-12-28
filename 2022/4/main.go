package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var fs, fe, ss, se int
		fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &fs, &fe, &ss, &se)
		if ss <= fe && se >= fe || fs <= se && fe >= ss {
			sum++
		}

	}
	fmt.Printf("%d\n", sum)
}
