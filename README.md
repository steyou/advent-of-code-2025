# Advent of Code 2025 (Golang)

Solutions for the 2025 edition of [Advent of Code](https://adventofcode.com/) (AoC)
implemented in Go 1.24.

## Usage

After cloning the repo and `cd`'ing into it:

```bash
go run . <input-file> <day> <part>
```

Example: solve Day 1 Part 2 using an input file stored at `inputs/day01.txt`.

```bash
go run . inputs/day01.txt 1 2
```

Puzzle inputs are not included in the repo (since sharing them is discouraged).
Access yours from the website.

**Note:** For day 2 part 2 specifically you will need to run `go generate` with your input
called `input/day2.txt` as it involves generating code ahead of time.

## Repository Layout

- `main.go` – argument parsing and day/part dispatch
- `dayX.go` – solutions for Day X Part 1 and Part 2
- `checkRepeating.go` - Placeholder file that is replaced for solving Day 2 Part 2

## Daily Solutions

### Day 1

[Link to file](https://github.com/steyou/advent-of-code-2025/blob/df39222cc16c7ea58af865c66a93d3f3b0ab9349/day1.go)

- **Part 1** Move the dial with modulo arithmetic, following Python's negative handling of the modulus operator rather than Go's native modulus. **O(n)** where n = number of commands.

- **Part 2** A variation of Part 1 with slightly harder math and handling an edge case if the dial is already at zero. **O(n)** where n = number of commands.

### Day 2

[Link to file](https://github.com/steyou/advent-of-code-2025/blob/e9af009759169341c66a7fb6611ab0209131795a/day2.go)

- **Part 1** Split even-length numbers into upper and lower halves using integer division and modulo by powers of 10, checking if they match. Skips odd-length numbers by jumping to next power of 10. **O(R)** where R = total range size across all input pairs.

- **Part 2** Generated a switch case ahead-of-time containing all numbers built with repeating patterns (~50k lines long file just on the sample input!). Actually faster than using the "correct" algorithm. **O(R)** runtime where R = total range size (code generation is O(M log M) where M = max number).

### Day 3

[Link to file](https://github.com/steyou/advent-of-code-2025/blob/b834784062e4db37926fa578a34bee2bc9db4dc6/day3.go)

- **Part 1** Single right-to-left pass tracking the maximum leading digit and the best trailing digit that can pair with it. **O(L·n)** where L = number of lines, n = digits per line.

- **Part 2** Greedy selection filling 12 positions left-to-right. For each position, scan remaining digits right-to-left to find the maximum digit that doesn't violate position constraints. **O(L·n)** where L = number of lines, n = digits per line (with constant factor for 12 slots).

### Day 4

[Link to file](https://github.com/steyou/advent-of-code-2025/blob/b834784062e4db37926fa578a34bee2bc9db4dc6/day4.go)

- **Part 1** Check each position in the grid by summing ASCII values of neighbors. Use precomputed threshold values (based on character ASCII codes) to determine if a position qualifies. Handle middle, edges, and corners separately. **O(w·h)** where w = width, h = height of grid.

- **Part 2** Same algorithm as Part 1, but remove qualifying positions from the grid after each pass and repeat until no more positions qualify. **O(k·w·h)** where k = number of iterations until convergence.

### Day 5

[Link to file](https://github.com/steyou/advent-of-code-2025/blob/4736ab6989550ac1d59b786cc3a357956301f89f/day5.go)

- **Part 1** Sort the ranges by starting point. For each query, binary search to find the rightmost range where `start <= x`, then scan backwards checking if `x` falls within any range. **O(n log n + k log n)** average case where n = number of ranges, k = number of queries. Worst case **O(kn)** if all ranges overlap.

- **Part 2** Create two independent arrays for range starts and ends, sort them separately. Use two pointers to process start/end events in order, tracking active range count. When the count drops to zero after decrementing, close the current merged range. **O(n log n)** where n = number of ranges.

### Day 6

[Link to file](https://github.com/steyou/advent-of-code-2025/blob/4736ab6989550ac1d59b786cc3a357956301f89f/day6.go)

- **Part 1** Read file backwards so that you can access the operators first and initialize an array of results with 1 or 0 depending on if the operator is multiply or add. Then parse every line and do the operation. **O(L·c)** where L = number of lines, c = columns per line.

- **Part 2** Read each column top-to-bottom, collecting non-space digits to form numbers. Parse the operator line into column ranges, then use binary search to map each number's column position to its corresponding operator for accumulation. **O(w·h + m log p)** where w = width, h = height of grid, m = numbers extracted, p = number of problems/operators.

### Day 7

[Link to file](https://github.com/steyou/advent-of-code-2025/blob/8b4070fcd60b43d2e1a4a5e1fb6775159976f2c4/day7.go)

- **Part 1** Just the obvious solution implied by the problem. Literally draw the beams in memory and count when you split.

- **Part 2** Store values over the width of the input, processing from the bottom up. Every time there is a diagonal `^`, increment the counter above that `^`. The answer is the value in the middle at the top of the input. This I find easier than trying to solve recursively.

### Day 8

[Link to file](https://github.com/steyou/advent-of-code-2025/blob/8b4070fcd60b43d2e1a4a5e1fb6775159976f2c4/day8.go)

- **Part 1** A modified Kruskall's algorithm.

- **Part 2** Modify the driver loop to observe the number of circuits and stop there.

### Day 9

[Link to file](https://github.com/steyou/advent-of-code-2025/blob/8b4070fcd60b43d2e1a4a5e1fb6775159976f2c4/day9.go)

- **Part 1** Brute-force measuring the area of all coordinates.

- **Part 2** Simply put, determine if all corners (ie including the two implied corners) are within the bounds of the polygon. After that check the edges don't intersect via bounds checking.

### Day 10

[Link to file](https://github.com/steyou/advent-of-code-2025/blob/8b4070fcd60b43d2e1a4a5e1fb6775159976f2c4/day10.go)

- **Part 1** If you represent bits as nodes and operations as edges you have a directed graph which you can traverse via recursive DFS. You keep a map to avoid duplicate states and have a means to terminate the recursion.

- **Part 2** This one goes over my head. I feel it involves a system of linear equations but I'm unsure about how to form them.

### Day 11

[Link to file](https://github.com/steyou/advent-of-code-2025/blob/8b4070fcd60b43d2e1a4a5e1fb6775159976f2c4/day11.go)

- **Part 1** Parse the graph twice: a first pass counts nodes and max fan-out, and a second pass maps names to dense indices while writing neighbor indices into a flat array. A memoized DFS (`countPaths`) then walks the DAG from `svr` to `out`, so runtime is **O(V+E)**.

- **Part 2** Uses the identical machinery except you run the program with multiple passes targeting `svr -> fft`, `fft -> dac`, `dac -> out`, and multiplying the outputs with a calculator.

### Day 12

[Link to file](https://github.com/steyou/advent-of-code-2025/blob/8b4070fcd60b43d2e1a4a5e1fb6775159976f2c4/day12.go)

- **Part 1** This is an NP-complete problem in the general case. Turns out the answer is the dumb one: just sum the areas. I guess there exists a way to tile the input perfectly. Here I've skipped parsing the top half - you need to provide a truncated input and provide the areas in the code.

- **Part 2** Cannot solve until Day 10 Part 2 is solved.
