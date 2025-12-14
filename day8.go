package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point8 struct {
	x, y, z, circuit int
}

type Edge8 struct {
	lhs, rhs int
	distance int // I don't actually need to square root since I'm comparing distances
}

type unionFind8 struct {
	parent []int
	rank   []int
}

func newUnionFind8(n int) *unionFind8 {
	parent := make([]int, n)
	rank := make([]int, n)
	for i := range parent {
		parent[i] = i
	}
	return &unionFind8{
		parent: parent,
		rank:   rank,
	}
}

func (uf *unionFind8) find(x int) int {
	if uf.parent[x] != x {
		uf.parent[x] = uf.find(uf.parent[x])
	}
	return uf.parent[x]
}

func (uf *unionFind8) union(a, b int) bool {
	rootA := uf.find(a)
	rootB := uf.find(b)
	if rootA == rootB {
		return false
	}
	if uf.rank[rootA] < uf.rank[rootB] {
		uf.parent[rootA] = rootB
	} else if uf.rank[rootA] > uf.rank[rootB] {
		uf.parent[rootB] = rootA
	} else {
		uf.parent[rootB] = rootA
		uf.rank[rootA]++
	}
	return true
}

func day8a(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	points := make([]Point8, 0, 1000)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		values := strings.Split(line, ",")

		x, err := strconv.Atoi(values[0])
		if err != nil {
			return err
		}
		y, err := strconv.Atoi(values[1])
		if err != nil {
			return err
		}
		z, err := strconv.Atoi(values[2])
		if err != nil {
			return err
		}

		points = append(points, Point8{
			x:       x,
			y:       y,
			z:       z,
			circuit: -1,
		})
	}

	n := len(points)
	edges := make([]Edge8, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			pointOuter := points[i]
			pointInner := points[j]
			x := pointOuter.x - pointInner.x
			y := pointOuter.y - pointInner.y
			z := pointOuter.z - pointInner.z
			dist := x*x + y*y + z*z
			edges = append(edges, Edge8{
				lhs:      i,
				rhs:      j,
				distance: dist,
			})
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		return edges[i].distance < edges[j].distance
	})

	uf := newUnionFind8(n)
	selected := 0

	circuits := make([]int, 0, 300)
	for _, e := range edges {
		if uf.union(e.lhs, e.rhs) {
			lhs := &points[e.lhs]
			rhs := &points[e.rhs]
			if lhs.circuit != -1 && rhs.circuit == -1 {
				rhs.circuit = lhs.circuit
				circuits[lhs.circuit]++
			} else if lhs.circuit == -1 && rhs.circuit != -1 {
				lhs.circuit = rhs.circuit
				circuits[rhs.circuit]++
			} else if lhs.circuit == -1 && rhs.circuit == -1 {
				circuits = append(circuits, 2)
				lhs.circuit = len(circuits) - 1
				rhs.circuit = lhs.circuit
			} else if lhs.circuit != rhs.circuit {
				fmt.Println("huh?")
			}
			selected++
			if selected == n-1 {
				break
			}
		}
	}

	sort.Ints(circuits)

	// fmt.Println(maxCircuit)
	fmt.Println(len(circuits))
	fmt.Println(circuits)

	return nil
}
