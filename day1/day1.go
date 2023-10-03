package day1

import (
	"bufio"
	"embed"
	"fmt"
	"strconv"
)

//go:embed input.txt

var f embed.FS

func readInput() ([][]int, error) {
	file, err := f.Open("input.txt")
	if err != nil {
		return nil, fmt.Errorf("reading input file: %w", err)
	}
	defer file.Close()

	fs := bufio.NewScanner(file)
	fs.Split(bufio.ScanLines)

	var input [][]int
	var current []int
	for fs.Scan() {
		l := fs.Text()
		if l == "" {
			input = append(input, current)
			current = nil
			continue
		}
		v, err := strconv.Atoi(l)
		if err != nil {
			return nil, fmt.Errorf("parsing line with value %q: %w", l, err)
		}
		current = append(current, v)
	}

	return input, nil
}

func TotalCaloriesBiggestElf() (int, error) {
	input, err := readInput()
	if err != nil {
		return 0, fmt.Errorf("reading input: %w", err)
	}

	var max int
	for _, elf := range input {
		var c int
		for _, v := range elf {
			c += v
		}
		if c > max {
			max = c
		}
	}

	return max, nil
}

func TotalTopThreeElves() ([3]int, error) {
	r := [3]int{}

	input, err := readInput()
	if err != nil {
		return r, fmt.Errorf("reading input: %w", err)
	}

	var idxNotAssigned, min int
	for _, elf := range input {
		var c int
		for _, v := range elf {
			c += v
		}
		if idxNotAssigned < 3 {
			r[idxNotAssigned] = c
			if min == 0 || c < min {
				min = c
			}
			idxNotAssigned++
			continue
		}
		if c <= min {
			continue
		}
		var newMin int
		for i := range r {
			if r[i] == min {
				r[i] = c
			}
			if newMin == 0 || r[i] < newMin {
				newMin = r[i]
			}
		}
		min = newMin
	}

	return r, nil
}
