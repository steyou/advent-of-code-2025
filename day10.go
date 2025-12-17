package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func searchToggles(state, target, depth uint, best *uint, buttons []uint, seen map[uint]uint) {
	for _, mask := range buttons {
		next := state ^ mask
		nextDepth := depth + 1
		if nextDepth >= *best {
			continue
		}
		if prevDepth, exists := seen[next]; exists && nextDepth >= prevDepth {
			continue
		}
		seen[next] = nextDepth
		if next == target {
			*best = nextDepth
			continue
		}
		searchToggles(next, target, nextDepth, best, buttons, seen)
	}
}

func day10a(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	var sum uint = 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		var target uint
		for i := 0; i < len(parts[0])-2; i++ {
			if parts[0][i+1] == '#' {
				target |= 1 << i
			}
		}

		buttons := make([]uint, len(parts)-2)
		for i := 1; i < len(parts)-1; i++ {
			bits := strings.Join(strings.Split(parts[i][1:], ","), "")
			for j := 0; j < len(bits)-1; j++ {
				buttons[i-1] |= 1 << (bits[j] - '0')
			}
		}

		bestDepths := make(map[uint]uint)
		bestDepths[0] = 0

		best := ^uint(0)
		searchToggles(0, target, 0, &best, buttons, bestDepths)
		sum += best
	}
	fmt.Println(sum)
	return nil
}
