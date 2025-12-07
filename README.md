# Advent of Code 2025 (Go)

Solutions for the 2025 edition of [Advent of Code](https://adventofcode.com/)
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

**Note:** For day 2 part B specifically you will need to run `go generate` with your input
called `input/day2.txt` as it involves generating code ahead of time.

## Repository Layout

- `main.go` – argument parsing and day/part dispatch
- `dayX.go` – solutions for Day X Part 1 and Part 2

## Progress (spoilers!)

### Day 1

[Link to file](https://github.com/steyou/advent-of-code-2025/blob/df39222cc16c7ea58af865c66a93d3f3b0ab9349/day1.go)

- **Part 1** Move the dial with modulo arithmetic, following Python's negative handling of the modulus operator rather than Go's native modulus.
- **Part 2** A variation of Part 1 with slightly harder math and handling an edge case if the dial is already at zero.

### Day 2

[Link to file](https://github.com/steyou/advent-of-code-2025/blob/e9af009759169341c66a7fb6611ab0209131795a/day2.go)

- **Part 1** "Masked" and compared the upper and lower half since only numbers of even length can have matches.

- **Part 2** Generated a switch case ahead-of-time containing all numbers built with repeating patterns (~50k lines long file just on the sample input!). Actually faster than using the "correct" algorithm.

### Day 3

[Link to file](https://github.com/steyou/advent-of-code-2025/blob/b834784062e4db37926fa578a34bee2bc9db4dc6/day3.go)

- **Part 1** Sweep from the right keeping the best tens digit plus any trailing digit that survives to its right.

- **Part 2** Walk the 12 slots left-to-right while a nested right-to-left window locks in each pick.

### Day 4

[Link to file](https://github.com/steyou/advent-of-code-2025/blob/b834784062e4db37926fa578a34bee2bc9db4dc6/day4.go)

- **Part 1** Do math on ASCII codes with various sliding windows for the middle and edges.

- **Part 2** The same except by modifying the grid in-memory after every pass and repeating until there were no changes.

### Day 5

[Link to file](https://github.com/steyou/advent-of-code-2025/blob/df39222cc16c7ea58af865c66a93d3f3b0ab9349/day1.go)

- **Part 1** Sort the ranges by starting point. Binary search to find a starting point for `x >= lower_range` then iterate to find ranges where `x` fits within them.

- **Part 2** Create two independent arrays for range starts and ends and sort them after parsing. Then increment/decrement a counter as you process opens and closes. If the counter equals zero after a decrement then create a new range.

### Day 6

[Link to file](https://github.com/steyou/advent-of-code-2025/blob/df39222cc16c7ea58af865c66a93d3f3b0ab9349/day1.go)

- **Part 1** Read file backwards so that you can access the operators first and initialize an array of results with 1 or 0 depending on if the operator is multiply or add. Then parse every line and do the operation.

- **Part 2** Theoretically you are repeating Part 1 except you apply a clockwise-then-vertical-mirror to the input first (dropping the last line of operators. A clockwise transformation is normally swapping `[r][c] -> [c][r]` followed by a horizonal flip. This code just does the former in memory. Mapping a column to an operation is done by storing column ranges in which a specific operator is used.
