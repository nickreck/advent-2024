package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part2() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewReader(file)

	safe := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		intSli := make([]int, 0)
		sli := strings.Fields(line)
		for _, v := range sli {
			intV, _ := strconv.Atoi(v)
			intSli = append(intSli, intV)
		}

		index := 0
		if checkIsSafe(&intSli, &index) {
			safe++
		}
	}
	fmt.Println(safe)
}

func checkIsSafe(sli *[]int, index *int) bool {
	test := make([]int, len(*sli))
	_ = copy(test, *sli)
	test = append(test[:*index], test[(*index)+1:]...)

	isDesc := sort.SliceIsSorted(test, func(x, y int) bool {
		return test[x] > test[y]
	})
	isAsc := sort.IntsAreSorted(test)

	if (isAsc || isDesc) && checkDistance(&test) {
		return true
	}

	(*index)++
	if *index == len(*sli) {
		return false
	}

	return checkIsSafe(sli, index)
}

func checkDistance(sli *[]int) bool {
	for i := 1; i < len(*sli); i++ {
		distance := abs((*sli)[i-1], (*sli)[i])
		if distance < 1 || distance > 3 {
			return false
		}
	}
	return true
}
