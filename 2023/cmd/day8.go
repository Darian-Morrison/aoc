package main

import (
	"fmt"
	"bufio"
	"os"
	"aoc"
	"strings"
)

type Pair struct {
	left string
	right string
}

func ParseInput(filename string) (string, map[string]Pair) {
	nodes := make(map[string]Pair)

	file := aoc.MustOpenFile(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	instructions := ""
	if scanner.Scan() {
		instructions = scanner.Text()
	}
	
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " = ")
		if len(parts) < 2 {
			continue
		}
		key := parts[0]
		optionsText :=  strings.Split(strings.Trim(parts[1], "()"), ", ")

		nodes[key] = Pair{
			left: optionsText[0],
			right: optionsText[1],
		}
	}
	return instructions, nodes
}

func AllTerminating(nodes []string) bool {
	for _, node := range nodes {
		if node[len(node) - 1] != 'Z' {
			return false
		}
	}
	return true
}

func InitialNodes(nodes map[string]Pair) []string {
	initialNodes := []string{}
	for node, _ := range nodes {
		if node[len(node) - 1] == 'A' {
			initialNodes = append(initialNodes, node)
		}
	}
	return initialNodes
}

func main(){
	args := os.Args[1:]
	
	inputPath := "cmd/day7.input"
	if len(args) < 1 {
		fmt.Printf("No input file path provided. Using default: %s\n", inputPath)
	} else {
		inputPath = args[0]
		fmt.Printf("Using %s as input\n", inputPath)
	}

	instructions, nodes := ParseInput(inputPath)	
	_, part1Compatible := nodes["ZZZ"]
	count := 0
	if part1Compatible {
		for index := "AAA"; index != "ZZZ"; count++ {
			if instructions[count % len(instructions)] == 'R' {
				index = nodes[index].right
			} else {
				index = nodes[index].left
			}
		}

		fmt.Printf("Number of steps (part1): %d\n", count)
	}

	count = 0
	startNodes := InitialNodes(nodes)
	for indexes := startNodes; !AllTerminating(indexes); count++ {
		newIndexes := []string{}
		if instructions[count % len(instructions)] == 'R' {
			for _, index := range indexes {
				newIndexes = append(newIndexes, nodes[index].right)
			}
		} else {
			for _, index := range indexes {
				newIndexes = append(newIndexes, nodes[index].left)
			}
		}
		fmt.Printf("Size of indexes: %d, count: %d\n", len(newIndexes), count)
		indexes = newIndexes
		
	}
	fmt.Printf("Number of steps (part2): %d\n", count)
}

