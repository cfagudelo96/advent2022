package main

import (
	"fmt"

	"github.com/cfagudelo96/advent2022/day2"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	in, err := day2.ReadInput(false)
	if err != nil {
		return fmt.Errorf("reading input for problem: %w", err)
	}
	r := day2.TotalCorrectScore(in)
	fmt.Println(r)
	return nil
}
