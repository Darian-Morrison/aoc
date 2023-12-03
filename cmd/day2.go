package main

import (
	"os"
	"strconv"
	"bufio"
	"fmt"
	"aoc"
	"strings"
)

func ParseGame(gameInput string) (int, []map[string]int) {
	parts := strings.Split(gameInput, ": ")
	title := parts[0]
	
	gameNumber, err := strconv.Atoi(title[5:])
	if err != nil {
		fmt.Printf("Error converting to integer: %v\n", err)
		return 0, []map[string]int{}
	}
	draws := strings.Split(parts[1], "; ")
	drawsParsed := make([]map[string]int, len(draws))
	for i, draw := range draws {
		drawsParsed[i] = map[string]int{
			"red": 0,
			"blue": 0,
			"green": 0,
		}
		for _, counts := range strings.Split(draw, ", ") {
			pair := strings.Split(counts, " ")	
			count, err := strconv.Atoi(pair[0])
			if err != nil {
				fmt.Printf("Error converting colour count: %v\n", err)
				return 0, []map[string]int{}
			}
			drawsParsed[i][pair[1]] += count
		}
	}
	return gameNumber, drawsParsed
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func FewestCubesPower(gameInput string) (int) {
	_, draws := ParseGame(gameInput)
	red := 0
	green := 0
	blue := 0
	for _, countsMap := range draws {
		red = max(red, countsMap["red"])	
		green = max(green, countsMap["green"])	
		blue = max(blue, countsMap["blue"])	
	}
	return red * green * blue
}

func ValidGame(gameInput string) (bool, int) {
	gameNumber, draws := ParseGame(gameInput)
	valid := true
	for _, countsMap := range draws {
		valid = valid && (countsMap["red"] <= 12 && countsMap["green"] <= 13 && countsMap["blue"] <= 14)
	}

	return valid, gameNumber
}


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

	sum := 0
	for scanner.Scan(){
		if part2 {
			sum += FewestCubesPower(scanner.Text())
		} else {
			valid, gameNumber := ValidGame(scanner.Text())
			if valid {
				sum += gameNumber
			}
		}
	}		
	fmt.Println(sum)
}

