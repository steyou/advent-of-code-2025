package main

import (
	"bufio"
	"fmt"
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

	scanner := bufio.NewScanner(file)

	// skip the top line
	scanner.Scan()

	// Read first data line
	scanner.Scan()
	lineBytes := scanner.Bytes()
	lineBytesLen := len(lineBytes)

	lines := make([]string, 0, int(fsize)/lineBytesLen)
	lines = append(lines, string(lineBytes))

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	linesLen := len(lines)
	values := make([]int, linesLen*lineBytesLen)

	// Initialize bottom row (base case)
	for i := 0; i < lineBytesLen; i++ {
		values[(linesLen-1)*lineBytesLen+i] = 1
	}

	// Bottom-up DP: work from second-to-last row upward
	for y := linesLen - 2; y >= 0; y-- {
		for x := 0; x < lineBytesLen; x++ {
			c := lines[y][x]
			if c == '^' {
				values[y*lineBytesLen+x] = values[(y+1)*lineBytesLen+(x-1)] + values[(y+1)*lineBytesLen+(x+1)]
			} else {
				values[y*lineBytesLen+x] = values[(y+1)*lineBytesLen+x]
			}
		}
	}

	// Print values as 2D grid for visualization
	for y := 0; y < linesLen; y++ {
		for x := 0; x < lineBytesLen; x++ {
			fmt.Printf("%3d ", values[y*lineBytesLen+x])
		}
		fmt.Println()
	}
	fmt.Println()

	// Find starting position and print result
	fmt.Println("Answer:", values[lineBytesLen>>1])
	return nil
}
