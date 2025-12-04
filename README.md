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

## Progress

| Day | Notes/hint |
| --- | --- |
| [Day 1](https://github.com/steyou/advent-of-code-2025/blob/df39222cc16c7ea58af865c66a93d3f3b0ab9349/day1.go) | Implemented using modulo math. |
| [Day 2](https://github.com/steyou/advent-of-code-2025/blob/e9af009759169341c66a7fb6611ab0209131795a/day2.go) | <p>**Part 1** "Masked" and compared the upper and lower half since only numbers of even length can have matches.</p><p>**Part 2** Generated a switch case ahead-of-time containing all numbers built with repeating patterns (~50k lines long file just on the sample input!). Actually faster than using the "correct" algorithm.</p>|
| [Day 3](https://github.com/steyou/advent-of-code-2025/blob/b834784062e4db37926fa578a34bee2bc9db4dc6/day3.go) | <p>**Part 1** Sweep from the right keeping the best tens digit plus any trailing digit that survives to its right.</p><p>**Part 2** Walk the 12 slots left-to-right while a nested right-to-left window locks in each pick.</p> |
| [Day 4](https://github.com/steyou/advent-of-code-2025/blob/b834784062e4db37926fa578a34bee2bc9db4dc6/day4.go) | Do math on ASCII codes with various sliding windows for the middle and edges. Part 2 is the same except by modifying the grid in-memory after every pass. |
