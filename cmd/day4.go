package main

import (
	"os"
	"fmt"
	"aoc"
	"bufio"
	"strings"
	"strconv"
)

func ScoreCard(game string) int {
	parts := strings.Split(game, " | ")
	winningNumbers := strings.Fields(parts[0])
	myNumbers := strings.Fields(parts[1])
	score := 0

	for _, winner := range winningNumbers {
		for _, num := range myNumbers {
			if(winner == num){
				if score == 0 {
					score = 1
				} else {
					score = score * 2
				}
				break
			}
		}
	}
	return score
}

type Stack struct {
	items[]interface{}
}

func (s *Stack) Push(item interface{}) {
	s.items = append(s.items, item)
}

func (s *Stack) Pop() interface{} {
	if s.IsEmpty() {
		return nil
	}
	top := s.items[len(s.items) - 1]
	s.items = s.items[:len(s.items)-1]
	return top
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func DecodeScore(score int, scoreMap map[int]int) int {
	value, exists := scoreMap[score]
	if exists{
		return value
	}
	count, acc := 1, score
	for acc > 1 {
		acc = acc / 2
		count++
	}
	scoreMap[score] = count
	return count
}

func main(){
	args := os.Args[1:]
	
	inputPath := "cmd/day4.input"
	if len(args) < 1 {
		fmt.Printf("No input file path provided. Using default: %s\n", inputPath)
	} else {
		inputPath = args[0]
		fmt.Printf("Using %s as input\n", inputPath)
	}
	file := aoc.MustOpenFile(inputPath)
	defer file.Close()
	
	scanner := bufio.NewScanner(file)

	sum := 0
	gameScore := make(map[int]int)	
	scoreMap := map[int]int{
		0: 0,
		1: 1,
	}
	stack := Stack{}
	count := 0

	for scanner.Scan(){
		parts := strings.Split(scanner.Text(), ": ")
		gameNum, err := strconv.Atoi(strings.TrimSpace(parts[0][5:]))
		if err != nil {
			fmt.Printf("Error converting game number: %v\n", err)
		}

		score := ScoreCard(parts[1])
		winningNumsCount := DecodeScore(score, scoreMap) 
		gameScore[gameNum] = score
		for i := gameNum + 1; i <= gameNum + winningNumsCount; i++ {
			stack.Push(i)
		}
		sum += gameScore[gameNum]
		count += 1
	}		
	fmt.Printf("Sum scratch card scores (part 1): %d\n", sum)
	for !stack.IsEmpty() {
		gameNum := stack.Pop().(int)
		score := gameScore[gameNum] 
		winningNumsCount := DecodeScore(score, scoreMap) 

		for i := gameNum + 1; i <= gameNum + winningNumsCount; i++ {
			stack.Push(i)
		}
		count += 1
	}

	fmt.Printf("Total number of scratchcards (part 2): %d\n", count)
}

