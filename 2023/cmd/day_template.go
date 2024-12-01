package main

import (
	"os"
	"strconv"
	"log"
	"bufio"
	"regexp"
	"fmt"
	"aoc"
	"unicode"
)


func main(){
	args := os.Args[1:]
	part2 := true
	if len(args) > 0 && args[0] == "--help" && ( args[0] != "--part=1" || args[0] != "--part=2") {
		fmt.Println("Valid options include --help, --part=1 and --part=2")
		os.Exit(0)
	}
	if len(args) > 0 && args[0] == "--part=1" {
		fmt.Println("Executing Part 1")
		part2 = false
	} else {
		fmt.Println("Executing Part 2")
	}
	
	file := aoc.MustOpenFile("cmd/day2.input")
	defer file.Close()
	
	scanner := bufio.NewScanner(file)

	fmt.Println("Finished!")
}

