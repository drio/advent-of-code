package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	top := []int{0, 0, 0, 0}
	sum := 0
	lineNum := 1

	updateIfBigger := func() {
		top[3] = sum
		sort.Sort(sort.Reverse(sort.IntSlice(top)))
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			updateIfBigger()
			sum = 0
			lineNum += 1
			continue
		}

		cals, err := strconv.Atoi(line)
		if err != nil {
			panic(err)
		}
		sum += cals
		lineNum += 1
	}

	updateIfBigger()
	fmt.Printf("%v %d\n", top[0:3], top[0]+top[1]+top[2])
}
