package main

import (
	"math"
	"bufio"
	"strings"
	"strconv"
	"io"
	"fmt"
	"os"
	"github.com/icza/backscanner"
)

type Problem6 struct {
	operator byte
	colStart, colEnd int
}

type Pair6 struct {
	foo, bar int
}

func day6a(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	fileStatus, err := file.Stat()
	if err != nil {
		return err
	}
	defer file.Close()
	// fmt.Println(fileStatus.Size())
	scanner := backscanner.New(file, int(fileStatus.Size()))

	// whoever wrote this made it so that the first invocation is an empty line
	// even if the file doesn't have one
	scanner.Line()

	line, _, err := scanner.Line()
	if err != nil {
		return err
	}

	ops := strings.Fields(line)

	accumulations := make([]int, len(ops))

	for i := 0; i < len(ops); i++ {
		if ops[i] == "*" {
			accumulations[i] = 1
		} else {
			accumulations[i] = 0
		}
	}

	for {
		line, _, err := scanner.Line()
		if err != nil {
			if err == io.EOF {
				break
			}
			return err
		}
		fields := strings.Fields(line)
		for i, field := range fields {
			f, err := strconv.Atoi(field)
			if err != nil {
				return err
			}
			if ops[i] == "*" {
				accumulations[i] *= f
			} else {
				accumulations[i] += f
			}
		}
	}
	sum := 0
	for _, acc := range accumulations {
		sum += acc
	}
	fmt.Println(sum)
	return nil
}

func day6b(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	fileStatus, err := file.Stat()
	if err != nil {
		return err
	}

	fsize := fileStatus.Size()
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	line := scanner.Text()
	lineLen := len(line)
	headerLines := int(fsize)/lineLen - 1

	grid := make([]string, headerLines)
	grid[0] = line

	for i := 1; i < headerLines; i++ {
		scanner.Scan()
		grid[i] = scanner.Text()
		if err != nil {
			break
		}
	}

	scanner.Scan()
	line = scanner.Text()

	var problems = []Problem6{}

	// Find each operator and determine its column range
	i := 0
	for i < len(line) {
		// Skip leading spaces
		for i < len(line) && line[i] == ' ' {
			i++
		}
		if i >= len(line) {
			break
		}

		// Found an operator
		rangeStart := i
		operator := line[i]

		// Find end of this field (either next non-space or end of line)
		i++
		for i < len(line) && line[i] == ' ' {
			i++
		}
		rangeEnd := i - 1

		problems = append(problems, Problem6{
			operator: operator,
			colStart: rangeStart,
			colEnd:   rangeEnd,
		})
	}

	// Initialize accumulations
	accumulations := make([]int, len(problems))
	for i := range problems {
		if problems[i].operator == '*' {
			accumulations[i] = 1
		} else {
			accumulations[i] = 0
		}
	}

	var nums = []Pair6{}

	digits := make([]byte, lineLen)
	for col := 0; col < lineLen; col++ {
		digitsLen := 0
		for row := 0; row < headerLines; row++ {
			x := grid[row][col]
			if x != ' ' {
				digits[digitsLen] = x - '0'
				digitsLen++
			}
		}
		if digitsLen > 0 {
			num := 0
			for e := 0; e < digitsLen; e++ {
				num += int(digits[e]) * int(math.Pow10(digitsLen - e - 1))
			}
			nums = append(nums, Pair6{
				foo:num,
				bar:col,
			})
		}
	}

	for _, pair := range nums {
		num := pair.foo
		col := pair.bar

		// Binary search for the problem containing this column
		left, right := 0, len(problems)-1
		for left <= right {
			mid := (left + right) / 2
			if col < problems[mid].colStart {
				right = mid - 1
			} else if col > problems[mid].colEnd {
				left = mid + 1
			} else {
				// Found it
				if problems[mid].operator == '*' {
					accumulations[mid] *= num
				} else {
					accumulations[mid] += num
				}
				break
			}
		}
	}

	sum := 0
	for _, acc := range accumulations {
		sum += acc
	}
	fmt.Println(sum)
	return nil
}
