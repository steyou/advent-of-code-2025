package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point9 struct {
	x, y int
}

func pointPairIndex9(n, i, j int) int {
	if i > j {
		i, j = j, i
	}
	return i*n - i*(i+1)/2 + j
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

	scanner := bufio.NewScanner(file)

	// Read all points and track max coordinate
	var points []Point9

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ",")
		x, err := strconv.Atoi(values[0])
		if err != nil {
			return err
		}
		y, err := strconv.Atoi(values[1])
		if err != nil {
			return err
		}
		points = append(points, Point9{
			x: x,
			y: y,
		})
	}

	// Calculate array size based on number of points
	n := len(points)
	numAreas := n * (n + 1) / 2
	areas := make([]int, numAreas)

	// Compute areas for all point pairs
	for i := 0; i < len(points); i++ {
		p1 := points[i]
		for j := i + 1; j < len(points); j++ {
			p2 := points[j]
			area := (abs(p2.x-p1.x) + 1) * (abs(p2.y-p1.y) + 1)
			areas[pointPairIndex9(n, i, j)] = area
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

type Edge9 struct {
	a, b Point9
}

type PointLocation int

const (
	Outside PointLocation = iota
	Inside
	Boundary
)

type Polygon struct {
	Vertices []Point9
	HEdges   []Edge9
	VEdges   []Edge9
}

func NewRectilinearPolygon(points []Point9) *Polygon {
	vertexSet := make(map[Point9]struct{}, len(points))
	for _, p := range points {
		vertexSet[p] = struct{}{}
	}

	byX := make(map[int][]Point9)
	for p := range vertexSet {
		byX[p.x] = append(byX[p.x], p)
	}

	var vEdges []Edge9
	for _, list := range byX {
		sort.Slice(list, func(i, j int) bool { return list[i].y < list[j].y })
		for i := 0; i < len(list); i += 2 {
			a := list[i]
			b := list[i+1]
			vEdges = append(vEdges, Edge9{a: a, b: b})
		}
	}

	byY := make(map[int][]Point9)
	for p := range vertexSet {
		byY[p.y] = append(byY[p.y], p)
	}

	var hEdges []Edge9
	for _, list := range byY {
		sort.Slice(list, func(i, j int) bool { return list[i].x < list[j].x })
		for i := 0; i < len(list); i += 2 {
			a := list[i]
			b := list[i+1]
			hEdges = append(hEdges, Edge9{a: a, b: b})
		}
	}

	polyVertices := make([]Point9, 0, len(vertexSet))
	for p := range vertexSet {
		polyVertices = append(polyVertices, p)
	}

	return &Polygon{
		Vertices: polyVertices,
		HEdges:   hEdges,
		VEdges:   vEdges,
	}
}

func (p *Polygon) ContainsPoint(q Point9) PointLocation {
	for _, e := range p.HEdges {
		y := e.a.y
		x1 := e.a.x
		x2 := e.b.x
		if x1 > x2 {
			x1, x2 = x2, x1
		}
		if q.y == y && q.x >= x1 && q.x <= x2 {
			return Boundary
		}
	}

	for _, e := range p.VEdges {
		x := e.a.x
		y1 := e.a.y
		y2 := e.b.y
		if y1 > y2 {
			y1, y2 = y2, y1
		}
		if q.x == x && q.y >= y1 && q.y <= y2 {
			return Boundary
		}
	}

	count := 0
	for _, e := range p.VEdges {
		x0 := e.a.x
		y1 := e.a.y
		y2 := e.b.y
		if y1 > y2 {
			y1, y2 = y2, y1
		}

		if q.y >= y1 && q.y < y2 && x0 > q.x {
			count++
		}
	}

	if count&1 == 1 {
		return Inside
	}

	return Outside
}

func day9b(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	// Read all points and track max coordinate
	var points []Point9

	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ",")
		x, err := strconv.Atoi(values[0])
		if err != nil {
			return err
		}
		y, err := strconv.Atoi(values[1])
		if err != nil {
			return err
		}
		points = append(points, Point9{
			x: x,
			y: y,
		})
	}

	poly := NewRectilinearPolygon(points)
	maxArea := -1
	for i := 0; i < len(points); i++ {
		outerPoint := points[i]
		for j := 0; j < len(points); j++ {
			foo := true
			if j == i {
				continue
			}
			innerPoint := points[j]

			// Check if implied corners are inside the polygon
			if poly.ContainsPoint(Point9{x: outerPoint.x, y: innerPoint.y}) == Outside {
				continue
			}
			if poly.ContainsPoint(Point9{x: innerPoint.x, y: outerPoint.y}) == Outside {
				continue
			}

			// Compute rectangle bounds
			minX, maxX := outerPoint.x, innerPoint.x
			if minX > maxX {
				minX, maxX = maxX, minX
			}
			minY, maxY := outerPoint.y, innerPoint.y
			if minY > maxY {
				minY, maxY = maxY, minY
			}

			// Check if any vertical edges cut through the rectangle
			for _, e := range poly.VEdges {
				x := e.a.x
				if x <= minX || x >= maxX {
					continue
				}
				y1, y2 := e.a.y, e.b.y
				if y1 > y2 {
					y1, y2 = y2, y1
				}
				if y2 > minY && y1 < maxY {
					foo = false
					goto nextPair
				}
			}

			// Check if any horizontal edges cut through the rectangle
			for _, e := range poly.HEdges {
				y := e.a.y
				if y <= minY || y >= maxY {
					continue
				}
				x1, x2 := e.a.x, e.b.x
				if x1 > x2 {
					x1, x2 = x2, x1
				}
				if x2 > minX && x1 < maxX {
					foo = false
					goto nextPair
				}
			}

		nextPair:
			if foo {
				// Valid rectangle found
				area := (abs(innerPoint.x-outerPoint.x) + 1) * (abs(innerPoint.y-outerPoint.y) + 1)
				if area > maxArea {
					maxArea = area
				}
			}
		}
	}

	fmt.Println(maxArea)
	return nil
}
