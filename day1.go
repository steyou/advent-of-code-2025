package main

import (
	"os"
	"fmt"
	"bufio"
	"strconv"
)

const CLOCK = 100

func day1a(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var dial int = 50
	var clicks uint = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		magnitude, err := strconv.Atoi(line[1:])
		if err != nil {
			return err
		}

		if line[0] == 'L' {
			// This emulates Python's way of doing modulo, eg -2 % 100 = 98, not =-2
			dial = (dial - magnitude) % CLOCK
			if dial < 0 {
				dial += CLOCK
			}

		} else {
			// line[0] should == 'R' here
			dial = (dial + magnitude) % CLOCK
		}

		if dial == 0 {
			clicks++
		}
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Println(clicks)

	return nil
}

func day1b(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var dial int = 50
	var clicks int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		magnitude, err := strconv.Atoi(line[1:])
		if err != nil {
			return err
		}

		if line[0] == 'L' {
			wraps := (magnitude - dial + CLOCK) / CLOCK
			if dial == 0 {
				// We don't want to count wrapping around here since we already
				// counted landing on zero on the last iteration
				wraps--;
			}
			if wraps > 0 {
				// Indirectly this means that magnitude >= 100, in which you
				// need to add more clicks
				clicks += wraps
			}

			dial = (dial - magnitude) % CLOCK
			if dial < 0 {
				dial += CLOCK
			}

		} else {
			// line[0] should == 'R' here
			tmpDial := dial + magnitude
			dial = tmpDial % CLOCK
			clicks += tmpDial / CLOCK
		}

		// if dial < 0 {
		// 	fmt.Println("mishandled underflow: ", line)
		// }
	}
	if err := scanner.Err(); err != nil {
		return err
	}

	fmt.Println(clicks)

	return nil
}
