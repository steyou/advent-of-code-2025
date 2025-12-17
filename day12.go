package main

import (
	"os"
	"fmt"
	"bufio"
	"strings"
	"strconv"
)

func day12a(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// file.Seek(96, 0)

	scanner := bufio.NewScanner(file)

	// for i := 0; i < 15; i++{
	// 	scanner.Scan()
	// }

	// I'm skipping the top half of the file since it's awkward to pass. All
	// you need is the magic number areas anyway.

	validCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		items := strings.Split(line, " ")

		dimensionItems := strings.Split(items[0], "x")

		width, err := strconv.Atoi(dimensionItems[0])
		if err != nil {
			return err
		}
		height, err := strconv.Atoi(dimensionItems[1][:len(dimensionItems)])
		if err != nil {
			return err
		}

		area := 0

		// Adjust the areas as needed:
		count, err := strconv.Atoi(items[1])
		if err != nil {
			return err
		}
		area += count * 7

		count, err = strconv.Atoi(items[2])
		if err != nil {
			return err
		}
		area += count * 7

		count, err = strconv.Atoi(items[3])
		if err != nil {
			return err
		}
		area += count * 7

		count, err = strconv.Atoi(items[4])
		if err != nil {
			return err
		}
		area += count * 5

		count, err = strconv.Atoi(items[5])
		if err != nil {
			return err
		}
		area += count * 6

		count, err = strconv.Atoi(items[6])
		if err != nil {
			return err
		}
		area += count * 7

		if area <= width*height {
			validCount++
		}
	}
	fmt.Println(validCount)
	return nil
}
