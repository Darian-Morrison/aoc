package aoc

import (
	"os"
	"log"
	"fmt"
	"strconv"
)


func Min(a int, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func StringsToInts(strList []string) []int {
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

func MustOpenFile(filename string) (*os.File) {
	file, error := os.Open(filename)
	if error != nil {
		log.Fatal(fmt.Sprintf("Failed to open s", file))
	}
	return file
}

