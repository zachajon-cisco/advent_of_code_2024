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
		values := strings.Split(l, " ")

		ok := true
		var incr bool
		var decr bool
		var same bool

		for i := range values {
			if i == len(values)-1 {
				break
			}
			v1, _ := strconv.Atoi(values[i])
			v2, _ := strconv.Atoi(values[i+1])

			if v1 < v2 {
				incr = true
				if (v2 - v1) > 3 {
					ok = false
				}
			} else if v1 > v2 {
				decr = true
				if (v1 - v2) > 3 {
					ok = false
				}
			} else {
				same = true
			}
		}

		if decr && incr {
			ok = false
		}

		if same {
			ok = false
		}

		if ok {
			safeReports += 1
		}
	}

	fmt.Println(safeReports)
}
