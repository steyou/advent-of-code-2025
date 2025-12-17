package main

import (
	"bufio"
	"fmt"
	"os"
)

func readGrid(fileName string) ([]byte, int, int, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, 0, 0, err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var lines [][]byte

	for scanner.Scan() {
		b := scanner.Bytes()
		row := append([]byte(nil), b...)
		lines = append(lines, row)
	}
	if err := scanner.Err(); err != nil {
		return nil, 0, 0, err
	}
	if len(lines) == 0 {
		return []byte{}, 0, 0, nil
	}

	width := len(lines[0])
	height := len(lines)

	grid := make([]byte, width*height)
	for y := 0; y < height; y++ {
		copy(grid[y*width:(y+1)*width], lines[y])
	}

	return grid, width, height, nil
}

func two_d_to_index(x, y, width int) int {
	return y*width + x
}

func day4a(fileName string) error {
	grid, width, h, err := readGrid(fileName)
	if err != nil {
		return err
	}

	sum := 0

	// count the middle
	for y := 1; y < h-1; y++ {
		for x := 1; x < width-1; x++ {
			if grid[two_d_to_index(x, y, width)] == '@' {
				var kernel int = int(grid[two_d_to_index(x-1, y-1, width)])
				kernel += int(grid[two_d_to_index(x, y-1, width)])
				kernel += int(grid[two_d_to_index(x+1, y-1, width)])

				kernel += int(grid[two_d_to_index(x-1, y, width)])
				kernel += int(grid[two_d_to_index(x+1, y, width)])

				kernel += int(grid[two_d_to_index(x-1, y+1, width)])
				kernel += int(grid[two_d_to_index(x, y+1, width)])
				kernel += int(grid[two_d_to_index(x+1, y+1, width)])

				// How I derived the "magic" 440 using ASCII values:
				//
				// >>> 46*8
				// 368
				// >>> 46*7+64
				// 386
				// >>> 46*6+64*2
				// 404
				// >>> 46*5+64*3
				// 422
				// >>> 46*4+64*4
				// 440
				// >>> 46*3+64*5
				// 458
				// >>> 46*2+64*6
				// 476
				// >>> 46*1+64*7
				// 494
				// >>> 46*0+64*8
				// 512
				//
				// Here since we need less than 4 `@` the condition is < 440
				if kernel < 440 {
					sum += 1
				}
			}
		}
	}

	// count the top row
	for x := 1; x < width-1; x++ {
		if grid[two_d_to_index(x, 0, width)] == '@' {
			var kernel int = int(grid[two_d_to_index(x-1, 0, width)])
			kernel += int(grid[two_d_to_index(x+1, 0, width)])

			kernel += int(grid[two_d_to_index(x-1, 1, width)])
			kernel += int(grid[two_d_to_index(x, 1, width)])
			kernel += int(grid[two_d_to_index(x+1, 1, width)])

			if kernel < 302 {
				sum += 1
			}
		}
	}

	// count the bottom row
	for x := 1; x < width-1; x++ {
		if grid[two_d_to_index(x, h-1, width)] == '@' {
			var kernel int = int(grid[two_d_to_index(x-1, h-1, width)])
			kernel += int(grid[two_d_to_index(x+1, h-1, width)])

			kernel += int(grid[two_d_to_index(x-1, h-2, width)])
			kernel += int(grid[two_d_to_index(x, h-2, width)])
			kernel += int(grid[two_d_to_index(x+1, h-2, width)])

			if kernel < 302 {
				sum += 1
			}
		}
	}

	// count the left column
	for y := 1; y < h-1; y++ {
		if grid[two_d_to_index(0, y, width)] == '@' {
			var kernel int = int(grid[two_d_to_index(0, y-1, width)])
			kernel += int(grid[two_d_to_index(0, y+1, width)])

			kernel += int(grid[two_d_to_index(1, y-1, width)])
			kernel += int(grid[two_d_to_index(1, y, width)])
			kernel += int(grid[two_d_to_index(1, y+1, width)])

			if kernel < 302 {
				sum += 1
			}
		}
	}

	// count the right column
	for y := 1; y < h-1; y++ {
		if grid[two_d_to_index(width-1, y, width)] == '@' {
			var kernel int = int(grid[two_d_to_index(width-1, y-1, width)])
			kernel += int(grid[two_d_to_index(width-1, y+1, width)])

			kernel += int(grid[two_d_to_index(width-2, y-1, width)])
			kernel += int(grid[two_d_to_index(width-2, y, width)])
			kernel += int(grid[two_d_to_index(width-2, y+1, width)])

			if kernel < 302 {
				sum += 1
			}
		}
	}

	// Resolve corners
	//
	// You don't need a loop or even counting because there are three
	// adjacent spaces and we are checking for less than four. Under that
	// condition it becomes a tautology
	if grid[two_d_to_index(0, 0, width)] == '@' {
		sum += 1
	}
	if grid[two_d_to_index(width-1, 0, width)] == '@' {
		sum += 1
	}
	if grid[two_d_to_index(0, h-1, width)] == '@' {
		sum += 1
	}
	if grid[two_d_to_index(width-1, h-1, width)] == '@' {
		sum += 1
	}

	fmt.Println(sum)

	return nil
}

type Coordinate struct {
	x int
	y int
}

func day4b(fileName string) error {
	// This is the same algorithm but the main change being you need to modify
	// the grid after you iterate it. The implication is that you need to
	// remember the coordinates to return later to make the changes. Eagerly
	// changing would break the search algorithm

	grid, width, h, err := readGrid(fileName)
	if err != nil {
		return err
	}

	sum := 0

	coordinates := []Coordinate{}

	for {
		// count the middle
		for y := 1; y < h-1; y++ {
			for x := 1; x < width-1; x++ {
				if grid[two_d_to_index(x, y, width)] == '@' {
					var kernel int = int(grid[two_d_to_index(x-1, y-1, width)])
					kernel += int(grid[two_d_to_index(x, y-1, width)])
					kernel += int(grid[two_d_to_index(x+1, y-1, width)])

					kernel += int(grid[two_d_to_index(x-1, y, width)])
					kernel += int(grid[two_d_to_index(x+1, y, width)])

					kernel += int(grid[two_d_to_index(x-1, y+1, width)])
					kernel += int(grid[two_d_to_index(x, y+1, width)])
					kernel += int(grid[two_d_to_index(x+1, y+1, width)])

					// How I derived the "magic" 440 using ASCII values:
					//
					// >>> 46*8
					// 368
					// >>> 46*7+64
					// 386
					// >>> 46*6+64*2
					// 404
					// >>> 46*5+64*3
					// 422
					// >>> 46*4+64*4
					// 440
					// >>> 46*3+64*5
					// 458
					// >>> 46*2+64*6
					// 476
					// >>> 46*1+64*7
					// 494
					// >>> 46*0+64*8
					// 512
					//
					// Here since we need less than 4 `@` the condition is < 440
					if kernel < 440 {
						coordinates = append(coordinates, Coordinate{
							x: x,
							y: y,
						})
					}
				}
			}
		}

		// count the top row
		for x := 1; x < width-1; x++ {
			if grid[two_d_to_index(x, 0, width)] == '@' {
				var kernel int = int(grid[two_d_to_index(x-1, 0, width)])
				kernel += int(grid[two_d_to_index(x+1, 0, width)])

				kernel += int(grid[two_d_to_index(x-1, 1, width)])
				kernel += int(grid[two_d_to_index(x, 1, width)])
				kernel += int(grid[two_d_to_index(x+1, 1, width)])

				if kernel < 302 {
					coordinates = append(coordinates, Coordinate{
						x: x,
						y: 0,
					})
				}
			}
		}

		// count the bottom row
		for x := 1; x < width-1; x++ {
			if grid[two_d_to_index(x, h-1, width)] == '@' {
				var kernel int = int(grid[two_d_to_index(x-1, h-1, width)])
				kernel += int(grid[two_d_to_index(x+1, h-1, width)])

				kernel += int(grid[two_d_to_index(x-1, h-2, width)])
				kernel += int(grid[two_d_to_index(x, h-2, width)])
				kernel += int(grid[two_d_to_index(x+1, h-2, width)])

				if kernel < 302 {
					coordinates = append(coordinates, Coordinate{
						x: x,
						y: h - 1,
					})
				}
			}
		}

		// count the left column
		for y := 1; y < h-1; y++ {
			if grid[two_d_to_index(0, y, width)] == '@' {
				var kernel int = int(grid[two_d_to_index(0, y-1, width)])
				kernel += int(grid[two_d_to_index(0, y+1, width)])

				kernel += int(grid[two_d_to_index(1, y-1, width)])
				kernel += int(grid[two_d_to_index(1, y, width)])
				kernel += int(grid[two_d_to_index(1, y+1, width)])

				if kernel < 302 {
					coordinates = append(coordinates, Coordinate{
						x: 0,
						y: y,
					})
				}
			}
		}

		// count the right column
		for y := 1; y < h-1; y++ {
			if grid[two_d_to_index(width-1, y, width)] == '@' {
				var kernel int = int(grid[two_d_to_index(width-1, y-1, width)])
				kernel += int(grid[two_d_to_index(width-1, y+1, width)])

				kernel += int(grid[two_d_to_index(width-2, y-1, width)])
				kernel += int(grid[two_d_to_index(width-2, y, width)])
				kernel += int(grid[two_d_to_index(width-2, y+1, width)])

				if kernel < 302 {
					coordinates = append(coordinates, Coordinate{
						x: width - 1,
						y: y,
					})
				}
			}
		}

		// Resolve corners
		//
		// You don't need a loop or even counting because there are three
		// adjacent spaces and we are checking for less than four. Under that
		// condition it becomes a tautology
		if grid[two_d_to_index(0, 0, width)] == '@' {
			coordinates = append(coordinates, Coordinate{
				x: 0,
				y: 0,
			})
		}
		if grid[two_d_to_index(width-1, 0, width)] == '@' {
			coordinates = append(coordinates, Coordinate{
				x: width - 1,
				y: 0,
			})
		}
		if grid[two_d_to_index(0, h-1, width)] == '@' {
			coordinates = append(coordinates, Coordinate{
				x: 0,
				y: h - 1,
			})
		}
		if grid[two_d_to_index(width-1, h-1, width)] == '@' {
			coordinates = append(coordinates, Coordinate{
				x: width - 1,
				y: h - 1,
			})
		}

		lenCoords := len(coordinates)
		if lenCoords == 0 {
			break
		}

		sum += lenCoords

		// Iterate through each coordinate, remove the @
		for _, coord := range coordinates {
			grid[two_d_to_index(coord.x, coord.y, width)] = '.'
		}

		// reset the slice
		coordinates = coordinates[:0]
	}

	fmt.Println(sum)

	return nil
}
