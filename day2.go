package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"
	"math"
)

func day2a(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	lineByComma := strings.Split(line, ",")
	for _, pair := range lineByComma {
		pairByDash := strings.Split(pair, "-")

		lower, err := strconv.Atoi(pairByDash[0])
		if err != nil {
			return err
		}
		upper, err := strconv.Atoi(pairByDash[1])
		if err != nil {
			return err
		}

		for i := lower; i <= upper; {
			// Get the length of this number.
			// Using the string length of `lower` or `upper` won't work since
			// `i` changes
			numLengthInt := len(strconv.Itoa(i))

			if numLengthInt & 1 == 1 {
				// Skip to next power of 10 (next even-length range)
				nextPow := int(math.Pow10(numLengthInt))
				if nextPow <= upper {
					i = nextPow
				} else {
					break
				}
				continue
			}

			numLengthHalf := numLengthInt >> 1
			upperHalf := i / int(math.Pow10(numLengthHalf))
			lowerHalf := i % int(math.Pow10(numLengthHalf))

			if upperHalf == lowerHalf {
				sum += i
			}

			i++
		}
	}
	fmt.Println(sum)
	return nil
}

func day2bMeta(fileName string) error {
	// You need to modify your input to run something like
	// sed -i 's/-/,/g' day2.txt
	// then get the original input back via the website to actually solve it

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	max := 0

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	lineByComma := strings.Split(line, ",")
	for _, num := range lineByComma {
		n, _ := strconv.Atoi(num)
		if n > max {
			max = n
		}
	}

	// You might think of doing a log10 but I'd rather avoid floating point
	// and rounding/off-by-one errors
	numLength := len(strconv.Itoa(max))

	// Use a set to track unique repeating numbers
	repeatingSet := make(map[int]struct{})

	// For each pattern width (1, 2, 3, ... digits)
	for patternWidth := 1; patternWidth <= numLength/2; patternWidth++ {
		// For each possible pattern of that width
		minPattern := int(math.Pow10(patternWidth - 1))
		maxPattern := int(math.Pow10(patternWidth)) - 1

		for pattern := minPattern; pattern <= maxPattern; pattern++ {
			// Repeat this pattern to build numbers
			current := pattern
			multiplier := int(math.Pow10(patternWidth))

			// Keep repeating until we exceed max
			for current <= max {
				// We need at least 2 repetitions to be a "repeating" number
				if current > pattern {
					repeatingSet[current] = struct{}{}
				}

				// Append the pattern again
				current = current*multiplier + pattern
			}
		}
	}

	fmt.Printf("switch (x) {\n")

	// Print all unique repeating numbers
	for num := range repeatingSet {
		fmt.Printf("case %d:\n\treturn true\n", num)
	}

	fmt.Println("}\nreturn false")

	return nil
}

func day2b(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	line := scanner.Text()

	lineByComma := strings.Split(line, ",")
	for _, pair := range lineByComma {
		pairByDash := strings.Split(pair, "-")

		lower, err := strconv.Atoi(pairByDash[0])
		if err != nil {
			return err
		}
		upper, err := strconv.Atoi(pairByDash[1])
		if err != nil {
			return err
		}

		for i := lower; i <= upper; i++ {
			if checkRepeating(i) {
				sum += i
			}
		}
	}
	fmt.Println(sum)
	return nil
}
