package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
	"strings"
	"slices"
	"cmp"
)

type Point8 struct {
	x, y, z, circuit int
}

type Edge8 struct {
	lhs, rhs int
	distance int
}

func pointPairIndex8(n, i, j int) int {
	if i > j {
		i, j = j, i
	}
	return i*n - i*(i+1)/2 + j
}

func day8a(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	points := [1000]Point8{}
	n := 0

	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Split(line, ",")

		x, err := strconv.Atoi(values[0])
		if err != nil {
			return err
		}
		y, err := strconv.Atoi(values[1])
		if err != nil {
			return err
		}
		z, err := strconv.Atoi(values[1])
		if err != nil {
			return err
		}

		points[n] = Point8{
			x: x,
			y: y,
			z: z,
			circuit:-1,
		}
		n++
	}

	edgesCount := (n * (n + 1)) >> 1
	edges := make([]Edge8, edgesCount)
	for i := 0; i < n; i++ {
		pointOuter := points[i]
		for j := i + 1; j < n; j++ {
			pointInner := points[j]

			x := pointOuter.x - pointInner.x
			y := pointOuter.y - pointInner.y
			z := pointOuter.z - pointInner.z
			dist := x*x + y*y + z*z

			edges = append(edges, Edge8{
				.lhs = i,
				.rhs = j,
				.distance = dist,
			})
		}
	}

	slices.SortFunc(edges, func(a, b Edge8) int {
		return cmp.Compare(a.distance, b.distance)
	})

	return nil
}
