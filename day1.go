package main

import (
	"fmt"
	"bufio"
)

func day1a(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}

func day1b(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	return nil
}
