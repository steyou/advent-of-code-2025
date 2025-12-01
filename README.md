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

## Repository Layout

- `main.go` – argument parsing and day/part dispatch
- `dayX.go` – solutions for Day X Part 1 and Part 2

## Progress

| Day | Notes/hint |
| --- | --- |
| [Day 1](https://github.com/steyou/advent-of-code-2025/blob/df39222cc16c7ea58af865c66a93d3f3b0ab9349/day1.go) | Implemented using modulo math |
