package main

import (
	"slices"
	"cmp"
	"strings"
	"strconv"
	"os"
	"fmt"
	"sort"
	"bufio"
)

type Pair struct {
	foo int
	bar int
}

func day5a(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	pairs := []Pair{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		lineByDash := strings.Split(line, "-")

		lower, err := strconv.Atoi(lineByDash[0])
		if err != nil {
			return err
		}
		upper, err := strconv.Atoi(lineByDash[1])
		if err != nil {
			return err
		}

		pairs = append(pairs, Pair{
			foo:lower,
			bar:upper,
		})
	}

	slices.SortFunc(pairs, func(a, b Pair) int {
		return cmp.Compare(a.foo, b.foo)
	})

	sum := 0
	for scanner.Scan() {
		line := scanner.Text()
		num, err := strconv.Atoi(line)
		if err != nil {
			return err
		}

		k := -1
		lower := 0
		upper := len(pairs)

		for lower < upper {
			mid := (lower + upper) >> 1
			if pairs[mid].foo <= num {
				k = mid
				lower = mid + 1
			} else {
				upper = mid
			}
		}

		for k >= 0 && pairs[k].foo <= num {
			if num >= pairs[k].foo && num <= pairs[k].bar {
				sum += 1
				break
			}
			k--
		}
	}
	fmt.Println(sum)
	return nil
}

func day5b(fileName string) error {
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	starts := []int{}
	ends := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			break
		}

		lineByDash := strings.Split(line, "-")

		lower, err := strconv.Atoi(lineByDash[0])
		if err != nil {
			return err
		}
		upper, err := strconv.Atoi(lineByDash[1])
		if err != nil {
			return err
		}

		starts = append(starts, lower)
		ends = append(ends, upper)
	}

	sort.Ints(starts)
	sort.Ints(ends)

	i := 0
	j := 0
	count := 0
	activeStart := -1
	pairs := []Pair{}

	for i < len(starts) || j < len(ends) {
		// Process whichever comes next: a start or an end
		if i < len(starts) && (j >= len(ends) || starts[i] <= ends[j]) {
			// Process a start
			if count == 0 {
				activeStart = starts[i]
			}
			count++
			i++
		} else {
			// Process an end
			count--
			if count == 0 {
				pairs = append(pairs, Pair{
					foo: activeStart,
					bar: ends[j],
				})
			}
			j++
		}
	}

	sum := len(pairs)
	for _, p := range pairs {
		sum += p.bar - p.foo
	}
	fmt.Println(sum)
	return nil
}
