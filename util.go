package aoc

import (
	"os"
	"log"
	"fmt"
)

func MustOpenFile(filename string) (*os.File) {
	file, error := os.Open(filename)
	if error != nil {
		log.Fatal(fmt.Sprintf("Failed to open s", file))
	}
	return file
}

