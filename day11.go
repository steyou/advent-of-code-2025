package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

type Node11 struct {
	NeighborsStart, NeighborsEnd int
}

func countPaths(nodes []Node11, indices, memo []int, source, target int) int {
	if source == target {
		return 1
	}

	if memo[source] != -1 {
		return memo[source]
	}

	node := nodes[source]
	total := 0

	for i := node.NeighborsStart; i < node.NeighborsEnd; i++ {
		total += countPaths(nodes, indices, memo, indices[i], target)
	}

	memo[source] = total
	return total
}

func day11a(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	maxNeighbors := 0
	nodeCount := 0
	for scanner.Scan() {
		nodeCount++

		lineBytes := scanner.Bytes()

		neighborCount := (len(lineBytes) >> 2) - 1
		if neighborCount > maxNeighbors {
			maxNeighbors = neighborCount
		}
	}

	maxNodes := nodeCount + 1 // +1 for "out" node
	nodes := make([]Node11, maxNodes)
	neighborIndices := make([]int, maxNeighbors*maxNodes)

	// Map node name string to integer Index
	nameToIndex := make(map[string]int, nodeCount)

	// Parse graph structure
	file.Seek(0, io.SeekStart)
	scanner = bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")

		nodeName := parts[0][:len(parts[0])-1]

		var nodeIndex int
		if id, exists := nameToIndex[nodeName]; exists {
			nodeIndex = id
		} else {
			nodeIndex = len(nameToIndex)
			nameToIndex[nodeName] = nodeIndex
		}

		neighborsStart := nodeIndex * maxNeighbors

		for i := 1; i < len(parts); i++ {
			neighborName := parts[i]
			var neighborIndex int
			if id, exists := nameToIndex[neighborName]; exists {
				neighborIndex = id
			} else {
				neighborIndex = len(nameToIndex)
				nameToIndex[neighborName] = neighborIndex
			}
			neighborIndices[neighborsStart+i-1] = neighborIndex
		}

		nodes[nodeIndex] = Node11{
			NeighborsStart: neighborsStart,
			NeighborsEnd:   neighborsStart + len(parts) - 1,
		}
	}

	youIndex, _ := nameToIndex["svr"]
	outIndex, _ := nameToIndex["out"]

	// youIndex, _ := nameToIndex["svr"]
	// outIndex, _ := nameToIndex["fft"]

	// outIndex, _ := nameToIndex["dac"]
	// outIndex, _ := nameToIndex["out"]
	memo := make([]int, len(nodes))
	for i := range memo {
		memo[i] = -1
	}
	sum := countPaths(nodes, neighborIndices, memo, youIndex, outIndex)
	fmt.Println(sum)

	// // Cycle detection using DFS with coloring
	// // Needed to figure out what algorithm to use since the problem didn't say
	// const (
	// 	white = 0 // Unvisited
	// 	gray  = 1 // Currently visiting
	// 	black = 2 // Finished visiting
	// )
	//
	// color := make([]int, len(nodes))
	// hasCycle := false
	//
	// var dfs func(int)
	// dfs = func(nodeIndex int) {
	// 	color[nodeIndex] = gray
	//
	// 	node := nodes[nodeIndex]
	// 	for i := node.NeighborsStart; i < node.NeighborsEnd; i++ {
	// 		neighborIndex := neighborIndices[i]
	// 		if color[neighborIndex] == gray {
	// 			hasCycle = true // Back edge found
	// 			return
	// 		} else if color[neighborIndex] == white {
	// 			dfs(neighborIndex)
	// 		}
	// 	}
	//
	// 	color[nodeIndex] = black
	// }
	//
	// // Visit all nodes (handles disconnected components)
	// for i := 0; i < len(nodes); i++ {
	// 	if color[i] == white {
	// 		dfs(i)
	// 	}
	// }
	//
	// fmt.Println("Has cycle:", hasCycle)

	return nil
}
