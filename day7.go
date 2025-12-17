package main

import (
	"bufio"
	"fmt"
	"github.com/icza/backscanner"
	"os"
)

func day7a(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Skip the top row since it's redundant for the algorithm
	scanner.Scan()

	scanner.Scan()
	top := append([]byte(nil), scanner.Bytes()...)

	lineLen := len(top)
	top[lineLen>>1] = '|'

	sum := 0
	for scanner.Scan() {
		middle := scanner.Text()

		scanner.Scan()
		bottom := append([]byte(nil), scanner.Bytes()...)

		for i := 0; i < lineLen; i++ {
			// read the middle line for caret characters and handle from there
			c := middle[i]
			if c == '^' {
				if top[i] == '|' {
					bottom[i-1] = '|'
					bottom[i+1] = '|'
					sum += 1
				}
			} else if c == '.' && top[i] == '|' {
				bottom[i] = '|'
			}
		}

		top = bottom
	}
	fmt.Println(sum)
	return nil
}

func day7b(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	fileStatus, err := file.Stat()
	if err != nil {
		return err
	}

	fsize := fileStatus.Size()

	backScanner := backscanner.New(file, int(fsize))

	// Skip trailing empty line
	backScanner.Line()

	// Get first line to determine width
	line, _, err := backScanner.Line()
	if err != nil {
		return err
	}

	lineBytesLen := len(line)

	// Allocate two row buffers
	currentRow := make([]int, lineBytesLen)
	nextRow := make([]int, lineBytesLen)

	// Initialize bottom row (base case)
	for x := 0; x < lineBytesLen; x++ {
		currentRow[x] = 1
		fmt.Printf("%3d ", currentRow[x])
	}
	fmt.Println()

	// Process all rows backwards
	for {
		// Compute next row based on current and line
		for x := 0; x < lineBytesLen; x++ {
			if line[x] == '^' {
				nextRow[x] = currentRow[x-1] + currentRow[x+1]
			} else {
				nextRow[x] = currentRow[x]
			}
		}

		// Print processed row
		for x := 0; x < lineBytesLen; x++ {
			fmt.Printf("%3d ", nextRow[x])
		}
		fmt.Println()

		// Swap buffers
		currentRow, nextRow = nextRow, currentRow

		// Fetch next line for next iteration
		line, _, err = backScanner.Line()
		if err != nil {
			break
		}
	}

	fmt.Println()
	fmt.Println("Answer:", currentRow[lineBytesLen>>1])
	return nil
}
