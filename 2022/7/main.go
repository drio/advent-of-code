package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type node struct {
	name     string
	size     int
	children map[string]*node
	parent   *node
}

func computeSize(n *node) int {
	s := n.size
	for _, p := range n.children {
		s += computeSize(p)
	}
	return s
}

func main() {
	var currNode *node
	var allNodes []*node
	var root *node
	reNum := regexp.MustCompile(`\d+`)
	s := bufio.NewScanner(os.Stdin)

	for s.Scan() {
		line := strings.Fields(s.Text())

		if line[0] == "$" && line[1] == "ls" {
			continue
		}

		if line[0] == "$" && line[1] == "cd" && line[2] == ".." {
			currNode = currNode.parent
			continue
		}

		if line[0] == "$" && line[1] == "cd" && line[2] != ".." {
			newNode := &node{
				name:     line[2],
				size:     0,
				children: make(map[string]*node),
				parent:   currNode,
			}

			if currNode != nil {
				currNode.children[line[2]] = newNode
			}

			if newNode.name == "/" {
				root = newNode
			}
			currNode = newNode
			allNodes = append(allNodes, newNode)
			continue
		}

		if reNum.MatchString(line[0]) {
			s, _ := strconv.Atoi(line[0])
			currNode.size += s
		}
	}

	totalSpace := 70000000
	needSpace := 30000000
	smallest := totalSpace
	used := computeSize(root)
	for _, n := range allNodes {
		s := computeSize(n)
		free := totalSpace - (used - s)
		if free >= needSpace && s < smallest {
			smallest = s
		}
	}
	fmt.Printf("%d\n", smallest)
}
