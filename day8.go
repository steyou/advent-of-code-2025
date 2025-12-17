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
	x, y, z int
}

type Edge8 struct {
	lhs, rhs int
	distance int // I don't actually need to square root since I'm comparing distances
}

type unionFind8 struct {
	parent []int
	size   []int
}

func newUnionFind8(n int) *unionFind8 {
	parent := make([]int, n)
	size := make([]int, n)
	for i := range parent {
		parent[i] = i
		size[i] = 1
	}
	return &unionFind8{
		parent: parent,
		size:   size,
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
	if uf.size[rootA] < uf.size[rootB] {
		rootA, rootB = rootB, rootA
	}
	uf.parent[rootB] = rootA
	uf.size[rootA] += uf.size[rootB]
	return true
}

func (uf *unionFind8) componentSizes() []int {
	sizes := make([]int, 0, len(uf.parent))
	for i := range uf.parent {
		root := uf.find(i)
		if root == i {
			sizes = append(sizes, uf.size[i])
		}
	}
	return sizes
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

		values := strings.Split(line, ",")

		x, _ := strconv.Atoi(values[0])
		y, _ := strconv.Atoi(values[1])
		z, _ := strconv.Atoi(values[2])

		points = append(points, Point8{
			x: x,
			y: y,
			z: z,
		})
	}

	n := len(points)
	edges := make([]Edge8, 0, n*(n-1)/2)
	for i := 0; i < n; i++ {
		pointOuter := points[i]
		for j := i + 1; j < n; j++ {
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

	// for i := 0; true; i++ {
	for i := 0; i < 1000; i++ {
		edge := edges[i]
		uf.union(edge.lhs, edge.rhs)

		//// Part 2:
		// if i > 1000 {
		// 	circuits := uf.componentSizes()
		// 	if len(circuits) == 1 {
		// 		fmt.Println(i)
		// 		fmt.Println(points[edge.lhs])
		// 		fmt.Println(points[edge.rhs])
		// 		break
		// 	}
		// }
	}

	circuits := uf.componentSizes()
	sort.Ints(circuits)

	fmt.Println(len(circuits))
	fmt.Println(circuits)

	return nil
}
