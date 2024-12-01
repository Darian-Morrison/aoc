package main

import (
	"fmt"
	"os"
	"bufio"
	"aoc"
	"strings"
	"strconv"
	"log"
)

type Race struct {
	time int
	distance int
}
func JoinRaces(races []Race) Race{
	timeBuf := ""
	distanceBuf := ""
	for _, race := range races {
		timeBuf = timeBuf + strconv.Itoa(race.time)
		distanceBuf = distanceBuf + strconv.Itoa(race.distance)
	}
	
	time, err := strconv.Atoi(timeBuf)
	if err != nil {
		log.Fatal("Error converting time to Int: %v\n", err)
	}
	distance, err := strconv.Atoi(distanceBuf)
	if err != nil {
		log.Fatal("Error converting distance to Int: %v\n", err)
	}
	return Race{
		time: time,
		distance: distance,
	}
}

func ParseInput(filename string) []Race {
	result := []Race{}

	file := aoc.MustOpenFile(filename)
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	parts := make(map[string][]int)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), ":")
		parts[split[0]] = aoc.StringsToInts(strings.Fields(split[1]))
	}

	for i, time := range parts["Time"] {
		result = append(result, Race{
			time: time,
			distance: parts["Distance"][i],
		})
	}
	return result
}

func WinningPermutations(race Race) int{
	time := race.time

	for i := 0; i <= time/2; i++ {
		distance := (time - i)*i
		if distance > race.distance {
			return (time + 1 - i) - i 
		}
	}
	return 0
}

func main(){
	args := os.Args[1:]
	
	inputPath := "cmd/day6.input"
	if len(args) < 1 {
		fmt.Printf("No input file path provided. Using default: %s\n", inputPath)
	} else {
		inputPath = args[0]
		fmt.Printf("Using %s as input\n", inputPath)
	}

	races := ParseInput(inputPath)
	total := 0
	for _, race := range races {
		// fmt.Printf("    Time: %d, Distance: %d\n", race.time, race.distance)
		ways := WinningPermutations(race)
		if total == 0 {
			total = ways
		} else {
			total = ways * total
		}
	}
	fmt.Printf("Product of winning permutations: %d\n", total)
	part2 := WinningPermutations(JoinRaces(races))
	fmt.Printf("Winning permutations (part 2): %d\n", part2)
}

