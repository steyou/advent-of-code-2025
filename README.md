# Advent of Code 2025 (Go)

Solutions for the 2025 edition of [Advent of Code](https://adventofcode.com/)
implemented in Go 1.24. Each puzzle day lives in its own file (e.g. `day1.go`)
and the `main` package dispatches to the requested day/part.

## Requirements

- Go 1.24 or newer (`go version` shows `go1.24.4` in this workspace)

## Usage

After cloning the repo and `cd`ing into it:

```bash
go run . <input-file> <day> <part>
```

Example: solve Day 1 Part 2 using an input file stored at `inputs/day01.txt`.

```bash
go run . inputs/day01.txt 1 2
```

My inputs are not provided. Get yours from the website.

## Repository Layout

- `main.go` – argument parsing and day/part dispatch
- `dayX.go` – solutions for Day X Part 1 (`dayXa`) and Part 2 (`dayXb`)
