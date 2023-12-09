package main

import (
	"fmt"
	"bufio"
	"os"
	"aoc"
	"strings"
	"strconv"
	"log"
)

func ParseInput(filename string) []Hand {
	result := []Hand{}

	file := aoc.MustOpenFile(filename)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	
	for scanner.Scan() {
		parts := strings.Fields(scanner.Text())
		bid, err := strconv.Atoi(parts[1])

		if err != nil {
			log.Fatal("Failed to convert bid to int: %v\n", err)
		}
		result = append(result, Hand{
			cards: parts[0],
			bid: bid,
		})
	}
	return result
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
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

	hands := SortHands(ParseInput(inputPath))
	size := len(hands)
	result := 0
	for i, hand := range hands {	
		result = result + (size - i) * hand.bid
	}
	fmt.Printf("Total winnings: %d\n", result)
}

