package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/zachajon-cisco/advent_of_code_2024/pkg/utils"
)

func main() {
	//A()
	B()

}

func B() {

	lines, err := utils.ReadInFile("day_1_a.txt")
	if err != nil {
		os.Exit(1)
	}

	var list1 []int
	var list2 []int

	for _, l := range lines {
		values := strings.Split(l, "   ")
		v, err := strconv.Atoi(values[0])
		if err != nil {
			os.Exit(1)
		}
		list1 = append(list1, v)
		v, err = strconv.Atoi(values[1])
		if err != nil {
			os.Exit(1)
		}
		list2 = append(list2, v)
	}

	//Sort lists
	sort.Ints(list1)
	sort.Ints(list2)

	//Determine difference
	var total int
	for i, _ := range list1 {
		a := list1[i]
		var count int
		for _, b := range list2 {
			if a == b {
				count += 1
			}
		}
		total += a * count
	}

	fmt.Println(total)
}

func A() {

	lines, err := utils.ReadInFile("day_1_a.txt")
	if err != nil {
		os.Exit(1)
	}

	var list1 []int
	var list2 []int

	for _, l := range lines {
		values := strings.Split(l, "   ")
		v, err := strconv.Atoi(values[0])
		if err != nil {
			os.Exit(1)
		}
		list1 = append(list1, v)
		v, err = strconv.Atoi(values[1])
		if err != nil {
			os.Exit(1)
		}
		list2 = append(list2, v)
	}

	//Sort lists
	sort.Ints(list1)
	sort.Ints(list2)

	//Determine difference
	var total int
	for i, _ := range list1 {
		a := list1[i] - list2[i]
		if a < 0 {
			a = a * -1
		}
		total += a
	}

	fmt.Println(total)
}
