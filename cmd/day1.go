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

var re_first = regexp.MustCompile(`(\d).*`)
var re_last = regexp.MustCompile(`.*(\d)`)
var numStringToDigit = map[string]rune{
	"one": '1',
	"two": '2',
	"three": '3',
	"four": '4',
	"five": '5',
	"six": '6',
	"seven": '7',
	"eight": '8',
	"nine": '9',
}

func DecodeLine(str string) int{
	first_match := re_first.FindStringSubmatch(str)
	last_match := re_last.FindStringSubmatch(str)

	num, err := strconv.Atoi(first_match[1] + last_match[1])
	if err != nil {
		log.Fatal("Error:", err)
	}
	return num
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func DecodeLine2(str string) int{
	var firstDigit, lastDigit rune
	
	for  i := 0; i < len(str); {
		runeValue := rune(str[i])
		if unicode.IsDigit(runeValue) {
			lastDigit = runeValue
			if firstDigit == 0 {
				firstDigit = runeValue
			}
		} else {
			for j, _ := range str[i:min(len(str), i + 6)] {
				digit, exists := numStringToDigit[str[i:(i+j+1)]]
				fmt.Printf("i: %d, i+j: %d, slice: %s, hit?: %t\n", i, i+j+1, str[i:(i+j+1)], exists)
				if exists {
					lastDigit = digit
					if firstDigit == 0 {
						firstDigit = digit
					}
					break
				}
					
			}
		}
		i++
	}
	fmt.Printf("Input: %s, first: %c, last: %c\n", str, firstDigit, lastDigit)
	num, err := strconv.Atoi(string(firstDigit) + string(lastDigit))
	if err != nil {
		log.Fatal("Error:", err)
	}
	return num
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
	
	file := aoc.MustOpenFile("cmd/day1.input")
	defer file.Close()
	
	scanner := bufio.NewScanner(file)

	sum := 0
	for scanner.Scan(){
		if(part2){
			sum += DecodeLine2(scanner.Text())
		} else {
			sum += DecodeLine(scanner.Text())
		}
	}		
	fmt.Println(sum)
}

