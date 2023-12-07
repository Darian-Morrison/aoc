package main

import (
	"os"
	"fmt"
	"aoc"
	"bufio"
	"strings"
	"strconv"
	"log"
)

type AlmanacMap struct{
	srcStart int
	srcEnd int
	destStart int
}

type AlmanacMapList struct{
	destination string
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

func BuildMap(filename string) map[string]AlmanacMapList {
	return nil			
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
	return seedNumbers, nil	
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
	seedNumbers, _ := ParseInput(inputPath)
	for _, seed := range seedNumbers {
		fmt.Printf("Seed: %d\n", seed)
	}
}

