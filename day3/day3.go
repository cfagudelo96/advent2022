package day3

import (
	"bufio"
	"embed"
	"fmt"
)

//go:embed input.txt
//go:embed testInput.txt

var f embed.FS

func ReadInput(test bool) ([]string, error) {
	path := "input.txt"
	if test {
		path = "testInput.txt"
	}
	file, err := f.Open(path)
	if err != nil {
		return nil, fmt.Errorf("reading input file %q: %w", path, err)
	}
	defer file.Close()

	fs := bufio.NewScanner(file)
	fs.Split(bufio.ScanLines)
	var input []string
	for fs.Scan() {
		input = append(input, fs.Text())
	}
	return input, nil
}

func SumOfPriorities(in []string) int {
	var sum int
	for _, s := range in {
		sum += linePriority(s)
	}
	return sum
}

func linePriority(s string) int {
	sbyte := []byte(s)
	middle := len([]byte(s)) / 2
	s1 := sbyte[:middle]
	s2 := sbyte[middle:]
	fmt.Printf("%v %v", s1, s2)
	return 0
}
