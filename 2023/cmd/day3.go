package main

import (
	"os"
	"fmt"
	"log"
	"io"
	"bufio"
	"unicode"
	"strconv"
)

type Pair struct{
	x int
	y int
}

func buildMatrix(path string) ([][]rune, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	 
	var result [][]rune	
	row := []rune{}
	for {
		r, _,  err := reader.ReadRune()
		if err == io.EOF {
			break // EOF
		} else if err != nil {
			return nil, err
		}

		if r == '\n' {
			result = append(result, row)
			row = []rune{}
		} else {
			row = append(row, r)
		}
	}

	if len(row) > 0 {
		result = append(result, row)
	}

	return result, nil
}

func IsAdjacentToSymbol(x, y int, matrix [][]rune) bool {
	// Check if out of range
	if y < 0 || y > len(matrix) - 1 || x < 0 || x > len(matrix[y]) - 1 {
		return false
	}
	r := matrix[y][x]
	if !unicode.IsDigit(r) &&  r != '.' {
		return true
	}
	return false
}

func IsPartNumber(x1, x2, y int, matrix[][]rune) bool{
	for i := x1; i <= x2; i++ {
		// Up
		up := IsAdjacentToSymbol(i, y - 1, matrix)
		// Down
		down := IsAdjacentToSymbol(i, y + 1, matrix)
		// Left
		left := IsAdjacentToSymbol(i - 1, y, matrix)
		// Right
		right := IsAdjacentToSymbol(i + 1, y, matrix)
		// Up, Right
		upRight := IsAdjacentToSymbol(i + 1, y - 1, matrix)
		// Down, Right
		downRight := IsAdjacentToSymbol(i + 1, y + 1, matrix)
		// Up, Left
		upLeft := IsAdjacentToSymbol(i - 1, y - 1, matrix)
		// Down, Left
		downLeft := IsAdjacentToSymbol(i - 1, y + 1, matrix)
			
		if up || down || left || right || upRight || downRight || upLeft || downLeft {
			return true
		}
	}
	return false
}

func SumOfPartNumbers(matrix [][]rune) int{
	sum := 0
	buf := ""

	for j, row := range matrix {
		for i, r := range row {

			if unicode.IsDigit(r) && i != 0 {
				buf += string(r)
				continue
			} 
			if len(buf) > 0 {
				// Check for part number if finished collecting digit string
				var x2, y int
				if i == 0 {
					// Handle case when going to next line
					y = j - 1
					x2 = len(matrix[j]) - 1
				} else {
					y = j
					x2 = i - 1
				}
				x1 := x2 - (len(buf) - 1)
				
				partNumber := IsPartNumber(x1, x2, y, matrix)
				if partNumber {
					num, err := strconv.Atoi(buf)
					if err != nil {
						log.Fatalf("Error converting buf to int: %v", err)
					}
					sum += num
				}
			} 
			buf = ""
			if unicode.IsDigit(r) && i == 0 {
				buf += string(r)
			}
			
		}
	}

	return sum
}

func FindNumber(x, y int, matrix [][]rune) int {
	if y < 0 || x < 0 || y >= len(matrix) || x >= len(matrix[y]) || !unicode.IsDigit(matrix[y][x]) {
		return -1
	}

	row := matrix[y]
	buf := string(row[x])
	i := 1
	var overflowLeft, overflowRight bool = false, false

	for {
		overflowLeft = overflowLeft || x - i < 0 || !unicode.IsDigit(row[x - i])
		overflowRight = overflowRight || x + i > len(row) - 1|| !unicode.IsDigit(row[x + i])

		if overflowLeft && overflowRight {
			break
		}
		if !overflowLeft {
			buf = string(row[x-i]) + buf
		}
		if !overflowRight {
			buf = buf + string(row[x + i])
		}
		i++
	}
	num, err := strconv.Atoi(buf)
	if err != nil {
		fmt.Printf("Error converting buffer to integer: %v\n", err)
	}
	return num
}

func CalculateGearRatio(symbol Pair, matrix [][]rune) int {
	var numbers []int

	// Left
	left := FindNumber(symbol.x - 1, symbol.y, matrix)
	if left != -1 {
		numbers = append(numbers,left) 
	}
	// Right
	right := FindNumber(symbol.x + 1, symbol.y, matrix)
	if right != -1 {
		numbers = append(numbers, right)
	}
	// Up
	up := FindNumber(symbol.x, symbol.y - 1, matrix)
	if(up == -1){
		// Up Left
		left := FindNumber(symbol.x - 1, symbol.y - 1, matrix)
		if left != -1 {
			numbers = append(numbers,left) 
		}
		// Up Right
		right := FindNumber(symbol.x + 1, symbol.y - 1, matrix)
		if right != -1 {
			numbers = append(numbers, right)
		}
	} else {
		numbers = append(numbers, up) 
	}
	// Down
	down := FindNumber(symbol.x, symbol.y  + 1, matrix)
	if(down == -1){
		// Down Left
		left := FindNumber(symbol.x - 1, symbol.y + 1, matrix)
		if left != -1 {
			numbers = append(numbers,left) 
		}
		// Down Right
		right := FindNumber(symbol.x + 1, symbol.y + 1, matrix)
		if right != -1 {
			numbers = append(numbers, right)
		}
	} else {
		numbers = append(numbers, down) 
	}

	if len(numbers) == 2 {
		return numbers[0] * numbers[1]
	}
	return 0
}

func SumOfGearRatios(matrix [][]rune) int {
	sum := 0
	
	symbols := []Pair{}
	// Locate Gear Symbols
	for j := 0; j < len(matrix); j++ {
		for i :=  0; i < len(matrix[j]); i++ {
			if matrix[j][i] == '*' {
				symbols = append(symbols, Pair{ x: i, y: j})
			}
		}
	}
	// Calculate ratio for each
	for _, symbol := range symbols {
		sum += CalculateGearRatio(symbol, matrix)	
	}
	return sum
}

func main(){
	args := os.Args[1:]
	
	inputPath := "cmd/day3.input"
	if len(args) < 1 {
		fmt.Printf("No input file path provided. Using default: %s\n", inputPath)
	} else {
		inputPath = args[0]
		fmt.Printf("Using %s as input\n", inputPath)
	}
	matrix, err := buildMatrix(inputPath)		
	if err != nil {
		log.Fatal("Error reading file: %v", err)
		return
	}
	
	fmt.Printf("Sum of part numbers: %d\n", SumOfPartNumbers(matrix))			
	fmt.Printf("Sum of gear ratio: %d\n", SumOfGearRatios(matrix))
}

