package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	argc := len(os.Args)
	if (argc != 4) {
		// fmt.Println("Expected day number. ie `go run . 1` for day 1")
		fmt.Printf("Expected: `%s file_path day part`\n", os.Args[0])
		return
	}

	dayNum, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	if !(dayNum >= 1 && dayNum <= 15) {
		fmt.Println("Invalid day, expected `1 <= arg <= 15`")
		return
	}

	partNum, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	if !(partNum >= 1 && partNum <= 3) {
		fmt.Println("Invalid part, expected 1 or 2")
		return
	}

	switch (dayNum) {
	case 1:
		if partNum == 1 {
			err = day1a(os.Args[1])
		} else {
			err = day1b(os.Args[1])
		}
	case 2:
		if partNum == 1 {
			err = day2a(os.Args[1])
		} else if partNum == 3 {
			err = day2bMeta(os.Args[1])
		} else {
			err = day2b(os.Args[1])
		}
	case 3:
		if partNum == 1 {
			err = day3a(os.Args[1])
		} else {
			err = day3b(os.Args[1])
		}
	case 4:
		if partNum == 1 {
			err = day4a(os.Args[1])
		} else {
			err = day4b(os.Args[1])
		}
	default:
		fmt.Println("fallback case")
		err = day1a(os.Args[1])
	}
	if err != nil {
		fmt.Println("error: ", err)
	}
}
