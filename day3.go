package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func day3a(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		var leadingMax byte = 0
		leadingPos := len(line) - 2

		trailingPos := len(line) - 1
		var trailingMax byte = line[trailingPos] - '0'

		trailingMaxAhead := trailingMax

		for i := leadingPos; i >= 0; i-- {
			num := line[i] - '0'
			// >= because even if == you still want to update the trailing
			if num >= leadingMax {
				if trailingMax < leadingMax {
					trailingMax = leadingMax
				}
				if trailingMax < trailingMaxAhead {
					trailingMax = trailingMaxAhead
				}
				leadingMax = num
				// leadingPos = i
			} else if num > trailingMax {
				trailingMaxAhead = num
			}
		}

		sum += int(leadingMax)*10 + int(trailingMax)
	}

	fmt.Println(sum)

	return nil
}

const sequenceLength = 12

func day3b(fileName string) error {
	// Unfortunately my code from Part A doesn't have much in common with how
	// to implement part B

	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lineLen := len(line)

		var sequence [sequenceLength]byte

		// Used to track where you must not seek past as the slot is taken.
		//
		// Initialized as -1 as you want to "disable" the condition to find the
		// absolute leftmost maximum number that later numbers anchor from
		tailStable := -1

		for i := 0; i < sequenceLength; i++ {
			tail := lineLen - sequenceLength + i

			sequence[i] = 0

			for c := tail; c >= 0 && c > tailStable; c-- {
				num := line[c] - '0'
				if num >= sequence[i] {
					sequence[i] = num
					tail = c
				}
			}

			tailStable = tail
		}

		for i := 0; i < sequenceLength; i++ {
			sum += int(sequence[i]) * int(math.Pow10(sequenceLength-i-1))
		}
		// if you wanted to show the chosen number for this line you'd print
		// sum - oldSum
	}
	fmt.Println(sum)
	return nil
}
