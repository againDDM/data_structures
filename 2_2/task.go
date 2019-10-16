package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type tree struct {
	name   int
	parent *tree
	childs []*tree
}

func readInput() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	line1, err := reader.ReadString('\n')
	if err != nil {
		panic("Can`t read")
	}
	line1 = line1[:len(line1)-1]
	line2, err := reader.ReadString('\n')
	if err != nil {
		panic("Can`t read")
	}
	line2 = line2[:len(line2)-1]
	return line1, line2
}

func parseInput(line1 string, line2 string) (numOfElem int, elemArray []*int) {
	numOfElem, _ = strconv.Atoi(line1)
	elemStringArray := strings.Split(line2, " ")
	for _, elemString := range elemStringArray {
		elemInt, _ := strconv.Atoi(elemString)
		elemArray = append(elemArray, &elemInt)
	}
	return
}

func fillNodes(nodes []*tree, nodeParentsId []*int) {
	for i, _ := range nodeParentsId {
		nodes[i] = new(tree)
	}

	for i, value := range nodeParentsId {
		nodes[i].name = i
		if *value != -1 {
			nodes[i].parent = nodes[*value]
			nodes[*value].childs = append(nodes[*value].childs, nodes[i])
		}
	}
}

func findRoot(nodes []*tree) (*tree, error) {
	for _, value := range nodes {
		if value.parent == nil {
			return value, nil
		}
	}
	return nil, fmt.Errorf("not found")
}

func getLength(treeToCheck *tree) int {
	var length int = 1
	for _, value := range treeToCheck.childs {
		length = max(length, 1+getLength(value))
	}
	return length
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func work(line1, line2 string) (int, error) {
	_, nodeParentsId := parseInput(line1, line2)
	nodes := make([]*tree, len(nodeParentsId))
	fillNodes(nodes, nodeParentsId)
	rootNode, err := findRoot(nodes)
	if err != nil {
		return 0, err
	}
	return getLength(rootNode), nil
}

func main() {
	output, err := work(readInput()); if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(output)
	}
}
