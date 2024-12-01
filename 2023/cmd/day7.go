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

type Hand struct {
	cards string
	bid int
}

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

func CountCards(cards string) map[rune]int {
	counter := make(map[rune]int)
	for _, card := range cards {
		count, exists := counter[card]
		if exists {
			counter[card] = count + 1
		} else {
			counter[card] = 0
		}
	}
	return counter
}

func CompareCardValues(a string, b string) int{
	options := "AKQJT9876543"
	for i := 0; i < len(a); i++ {
		var aRank, bRank = len(options), len(options)
		for j, _ := range options {
			if aRank != len(options) || bRank != len(options) {
				break
			}
			if options[j] == a[i]{
				aRank = j
			}
			if options[j] == b[i] {
				bRank = j
			}
		}
		if aRank < bRank {
			return 1
		}
		if aRank > bRank {
			return -1
		}
	}
	return 0
}

func CompareCards(a string, b string) int{
	aCounter := CountCards(a)
	bCounter := CountCards(b)
	// Lower size means bigger pair
	if len(aCounter) < len(bCounter){
		return 1
	} else if len(aCounter) > len(bCounter) {
		return -1
	}

	switch len(aCounter){
	case 2:
		// Handles Four of a kind and Full house
		fallthrough
	case 3:
		// Handles Three of a kind and two pair
		maxA := 0
		for _, value := range aCounter {
			maxA = aoc.Max(maxA, value)
		}
		maxB := 0
		for _, value := range bCounter {
			maxB = aoc.Max(maxB, value)	
		}
		if maxA > maxB {
			return 1
		} 
		if maxA < maxB {
			return -1
		}
		fallthrough
	default:
		// Handles Five of a kind, One pair and High card
		return CompareCardValues(a,b)
	}

	return 0
}

func SortHands(hands []Hand) []Hand {
	for i:=0; i < len(hands) - 1 ;i++ {
		swapped := false
		for j:=0; j < len(hands) - i - 1; j++ {
			result := CompareCards(hands[j].cards, hands[j + 1].cards)
			if result < 0 {
				temp := hands[j]
				hands[j] = hands[j + 1]
				hands[j + 1] = temp
				swapped = true
			}
		}

		if swapped == false {
			break
		}	
	}
	return hands
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

