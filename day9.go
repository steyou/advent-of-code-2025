package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point9 struct {
	x,y int
}

func coordPairIndex(n, ax, ay, bx, by int) int {
	p := n*ax+ay
	q := n*bx+by
	if p < q {
		p, q = q, p
	}
	return p*(p+1)/2 + q
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func day9a(fileName string) error {
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

	scanner.Scan()
	line := scanner.Text()

	suggestedPoints := int(fsize) / len(line)
	spBySp := suggestedPoints * suggestedPoints
	numAreas := spBySp*(spBySp+1)/2
	areas := make([]int, numAreas)

	points := make([]Point9, 0, suggestedPoints)

	values := strings.Split(scanner.Text(), ",")
	x1, err := strconv.Atoi(values[0])
	if err != nil {
		return err
	}
	y1, err := strconv.Atoi(values[1])
	if err != nil {
		return err
	}

	points = append(points, Point9{
		x:x1,
		y:y1,
	})

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ",")
		x2, err := strconv.Atoi(values[0])
		if err != nil {
			return err
		}
		y2, err := strconv.Atoi(values[1])
		if err != nil {
			return err
		}

		points = append(points, Point9{
			x:x2,
			y:y2,
		})

		areas[coordPairIndex(numAreas, x1, y1, x2, y2)] = abs(x2-x1) * abs(y2-y1)
	}

	for i := 1; i < len(points); i++ {
		outerPoint := points[i]
		for j := 0; j < len(points); j++ {
			innerPoint := points[j]
			areas[coordPairIndex(numAreas, outerPoint.x, outerPoint.y, innerPoint.x, innerPoint.y)] = abs(innerPoint.x-outerPoint.x) * abs(innerPoint.y-outerPoint.y)
		}
	}

	maxArea := -1
	for _, a := range areas {
		if a > maxArea {
			maxArea = a
		}
	}
	fmt.Println(maxArea)

	return nil
}
