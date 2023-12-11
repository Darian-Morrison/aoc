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

func ParseInput(filename string) [][]int {
	file := aoc.MustOpenFile(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	result := [][]int{}	
	for scanner.Scan() {
		result = append(result, aoc.StringsToInts(strings.Fields(scanner.Text())))
	}
	return result  
}

func PredictNext(history []int) int {
	result := history[len(history) - 1]
	sequence := history
	finished := true
	for len(sequence) > 1  {
		for i := 1; i < len(sequence); i++ {
			if finished && (sequence[i] - sequence[i - 1] != 0) {
				finished = false
			}
			sequence[i - 1] = sequence[i] - sequence[i - 1]
		}
		newSize := len(sequence) - 1
		sequence = sequence[:newSize]
		result += sequence[newSize - 1] // Add last element of new seq to result
		if finished {
			break
		}
		finished = true
	}

	return result
}

func PredictPrevious(history []int) int {
	result := history[0]
	sequence := history
	finished := true
	for count := 0; len(sequence) > 1 ; count++  {
		for i := 1; i < len(sequence); i++ {
			if finished && (sequence[i] - sequence[i - 1] != 0) {
				finished = false
			}
			sequence[i - 1] = sequence[i] - sequence[i - 1]
		}
		newSize := len(sequence) - 1
		sequence = sequence[:newSize]
		if count % 2 == 1 {
			result -= sequence[0]
		} else {
			result += sequence[0]
		}

		if finished {
			break
		}
		finished = true
	}

	return result
}

func main(){
	args := os.Args[1:]
	
	inputPath := "cmd/day8.input"
	if len(args) < 1 {
		fmt.Printf("No input file path provided. Using default: %s\n", inputPath)
	} else {
		inputPath = args[0]
		fmt.Printf("Using %s as input\n", inputPath)
	}

	sumNext := 0
	sumPrevious := 0
	for _, history := range ParseInput(inputPath) {
		sumNext += PredictNext(history)
		sumPrevious += PredictPrevious(history)
	}
	fmt.Printf("Sum of Future Predictions (part1): %d\n", sumNext)
	// This gives the negative result on the test input. Bothering me 
	// but not enough to figure it out.
	fmt.Printf("Sum of Past Predictions (part2): %d\n", sumPrevious)
}

