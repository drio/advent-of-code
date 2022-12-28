package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const debug = false

func main() {
	var field [][]int
	s := bufio.NewScanner(os.Stdin)

	// Load the field
	height := 0
	for s.Scan() {
		var row []int
		for _, c := range strings.Split(s.Text(), "") {
			ic, _ := strconv.Atoi(c)
			row = append(row, ic)
		}
		field = append(field, row)
		height += 1
	}
	width := len(field[0])

	bestScore := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			fmt.Printf("(%d,%d)='%d'\n ", x, y, field[x][y])
			score := scoreH(field, x, y, width) * scoreV(field, x, y, height)
			if score > bestScore {
				bestScore = score
			}
		}
	}
	fmt.Printf("%d\n", bestScore)
}

func scoreH(field [][]int, treeX, treeY, w int) int {
	treeH := field[treeX][treeY]

	countOne := 0
	for x := treeX - 1; x >= 0; x-- {
		countOne += 1
		if field[x][treeY] >= treeH {
			break
		}
	}

	countTwo := 0
	for x := treeX + 1; x < w; x++ {
		countTwo += 1
		if field[x][treeY] >= treeH {
			break
		}
	}
	return countOne * countTwo
}

func scoreV(field [][]int, treeX, treeY, h int) int {
	treeH := field[treeX][treeY]

	countOne := 0
	for y := treeY - 1; y >= 0; y-- {
		countOne += 1
		if field[treeX][y] >= treeH {
			break
		}
	}

	countTwo := 0
	for y := treeY + 1; y < h; y++ {
		countTwo += 1
		if field[treeX][y] >= treeH {
			break
		}
	}
	return countOne * countTwo
}

func log(msg string) {
	if debug {
		fmt.Printf("%s", msg)
	}
}
