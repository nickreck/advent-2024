package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewReader(file)

	safe := 0
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		sli := strings.Fields(line)
		intSli := make([]int, 0)
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

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}

func checkIsSafe(sli *[]int, index *int) bool {
	test := make([]int, len(*sli))
	_ = copy(test, *sli)
	test = append(test[:*index], test[(*index)+1:]...)

	isDesc := sort.SliceIsSorted(test, func(x, y int) bool {
		return test[x] > test[y]
	})
	isAsc := sort.IntsAreSorted(test)

	if isAsc || isDesc {
		for i := 1; i < len(test); i++ {
			distance := abs(test[i-1], test[i])
			if distance < 1 || distance > 3 {
				break
			}
			if i == len(test)-1 {
				return true
			}
		}
	}

	(*index)++
	if *index == len(*sli) {
		return false
	}

	return checkIsSafe(sli, index)
}
