package main

import (
	"math"
	"os"
	"fmt"
	"aoc"
	"bufio"
	"strings"
	"strconv"
	"log"
	"unicode"
)

type AlmanacMap struct{
	srcStart int
	srcEnd int
	destStart int
}

type AlmanacMapList struct{
	dest string
	maps []AlmanacMap
}

func (mapList *AlmanacMapList) FindDestination (input int) int{
	for _, almanacMap := range mapList.maps {
		if input <= almanacMap.srcEnd && input >= almanacMap.srcStart {
			return almanacMap.destStart + input - almanacMap.srcStart
		}
	}

	return input
}

func StringListToInts(strList []string) []int {
	var numList []int
	for _, str := range strList {
		num, err := strconv.Atoi(str)
		if err != nil{
			log.Fatal("Error converting num (%s) to int: %v\n", str, err)
		}
		numList = append(numList, num)
	}
	return numList
}

func BuildMap(scanner *bufio.Scanner) map[string]AlmanacMapList {
	almanac  := make(map[string]AlmanacMapList)
	var src, dest string

	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			continue
		} else if unicode.IsLetter(rune(line[0])) {
			// Create new AlmanacMapList
			srcDest := strings.Split(strings.Fields(line)[0], "-to-")
			src = srcDest[0]
			dest = srcDest[1]
			almanac[src] = AlmanacMapList{
				dest: dest,
				maps: []AlmanacMap{},
			}	
		} else if unicode.IsDigit(rune(line[0])) {
			// add new map list item
			mapFields := StringListToInts(strings.Fields(line))
			mapList := almanac[src]

			mapList.maps = append(almanac[src].maps, AlmanacMap{
				srcStart: mapFields[1],
				srcEnd: mapFields[1] + mapFields[2] - 1,
				destStart: mapFields[0],
			})
			almanac[src] = mapList
		}
	}
	return almanac			
}

func SeedToLocation(seed int, almanac map[string]AlmanacMapList) int {
	componentIndex := "seed"
	num := seed
	for componentIndex != "location" {
		mapList := almanac[componentIndex]
		for _, mapItem := range mapList.maps {
			if mapItem.srcStart <= num && mapItem.srcEnd >= num {
				num = mapItem.destStart + num - mapItem.srcStart
				break
			}
		}
		componentIndex = mapList.dest
	}
	return num
}

func ParseInput(filename string) ([]int,map[string]AlmanacMapList) {
	file := aoc.MustOpenFile(filename)
	defer file.Close()
	
	scanner := bufio.NewScanner(file)
	var seedNumbers []int

	if scanner.Scan() {
		seedStrings := strings.Fields(strings.Split(scanner.Text(), ": ")[1])
		
		for _, str := range seedStrings {
			seed, err := strconv.Atoi(str)
			if err != nil{
				log.Fatal("Error converting seed (%s) to int: %v\n", str, err)
			}
			seedNumbers = append(seedNumbers, seed)
		}
	}	
	return seedNumbers, BuildMap(scanner)	
}	

func main(){
	args := os.Args[1:]
	
	inputPath := "cmd/day5.input"
	if len(args) < 1 {
		fmt.Printf("No input file path provided. Using default: %s\n", inputPath)
	} else {
		inputPath = args[0]
		fmt.Printf("Using %s as input\n", inputPath)
	}
	seedNumbers, almanac := ParseInput(inputPath)
	minLocation := math.MaxInt

	for _, seed := range seedNumbers {
		newLocation := SeedToLocation(seed, almanac)
		if minLocation > newLocation {
			minLocation = newLocation
		}
	}
	fmt.Printf("Minimum location(part1): %d\n", minLocation)	
	minLocation = math.MaxInt	
	for i := 0; i < len(seedNumbers); i += 2 {
		start := seedNumbers[i]
		rng := seedNumbers[i + 1]
		for j := start; j < start + rng; j++ {
			newLocation := SeedToLocation(j, almanac)
			if minLocation > newLocation {
				minLocation = newLocation
			}
		}
	}
	fmt.Printf("Minimum location: %d\n", minLocation)	

}

