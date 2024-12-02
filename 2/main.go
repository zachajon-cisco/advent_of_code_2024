package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/zachajon-cisco/advent_of_code_2024/pkg/utils"
)

func main() {
	lines, err := utils.ReadInFile("day_2_a.txt")
	if err != nil {
		os.Exit(1)
	}

	var safeReports int

	for _, l := range lines {
		var problems int
		values := strings.Split(l, " ")
		for i := range values {

			if i == len(values)-1 {
				break
			}
			v1, _ := strconv.Atoi(values[i])
			v2, _ := strconv.Atoi(values[i+1])
			problems += SafeOrNot(v1, v2)
		}

		if problems <= 1 {
			safeReports += 1
		}

	}

	fmt.Println(safeReports)
}

func SafeOrNot(v1 int, v2 int) int {
	var incr bool
	var decr bool

	if v1 < v2 {
		incr = true
		if decr {
			return 1
		}
		if (v2 - v1) > 3 {
			return 1
		}
	} else if v1 > v2 {
		decr = true
		if incr {
			return 1
		}
		if (v1 - v2) > 3 {
			return 1
		}
	} else {
		return 1
	}

	return 0
}
