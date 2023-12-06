package main

import (
	"os"
	"fmt"
	"aoc"
	"bufio"
)

func main(){
	args := os.Args[1:]
	
	inputPath := "cmd/day5.input"
	if len(args) < 1 {
		fmt.Printf("No input file path provided. Using default: %s\n", inputPath)
	} else {
		inputPath = args[0]
		fmt.Printf("Using %s as input\n", inputPath)
	}
	file := aoc.MustOpenFile(inputPath)
	defer file.Close()
	
	scanner := bufio.NewScanner(file)

	count := 0	
	for scanner.Scan(){
		scanner.Text()	
		count++
	}		

	fmt.Printf("Number of lines: %d\n", count)
}

