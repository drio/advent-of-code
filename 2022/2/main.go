package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	op := map[string]string{"A": "R", "B": "P", "C": "S"}
	points := map[string]int{"R": 1, "P": 2, "S": 3}
	result := map[string]int{
		"RR": 3, "RP": 6, "RS": 0,
		"PR": 0, "PP": 3, "PS": 6,
		"SR": 6, "SP": 0, "SS": 3,
	}
	letterToOutcome := map[string]string{
		"X": "L", "Y": "D", "Z": "W",
	}
	ifPlay := map[string]string{
		"RD": "R", "RW": "P", "RL": "S",
		"PD": "P", "PW": "S", "PL": "R",
		"SD": "S", "SW": "R", "SL": "P",
	}

	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " ")
		opPlay := op[line[0]]
		out := letterToOutcome[line[1]]
		idxIf := fmt.Sprintf("%s%s", opPlay, out)
		myPlay := ifPlay[idxIf]
		idxResult := fmt.Sprintf("%s%s", opPlay, myPlay)
		r := result[idxResult] + points[myPlay]
		fmt.Printf("%s(%s) .(%s) %s(%s) %d+%d=%d %s\n",
			line[0], opPlay,
			myPlay,
			line[1], out,
			result[idxResult], points[myPlay], r,
			idxIf)
		sum += r
	}
	fmt.Printf("%d", sum)
}
