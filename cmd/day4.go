package main

import (
	"os"
	"fmt"
)

func main(){
	args := os.Args[1:]
	
	inputPath := "cmd/day3.input"
	if len(args) < 1 {
		fmt.Printf("No input file path provided. Using default: %s\n", inputPath)
	} else {
		inputPath = args[0]
		fmt.Printf("Using %s as input\n", inputPath)
	}
}

