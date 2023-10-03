package main

import (
	"fmt"

	"github.com/cfagudelo96/advent2022/day1"
)

func main() {
	r, err := day1.TotalTopThreeElves()
	if err != nil {
		panic(err)
	}
	fmt.Println(r)
}
