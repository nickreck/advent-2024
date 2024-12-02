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

		if !isSafe(&intSli) {
			for i := 0; i < len(intSli); i++ {
        test := make([]int, len(intSli))
        _ = copy(test, intSli)
				test = append(test[:i], test[i+1:]...)
				if isSafe(&test) {
					safe++
					break
				}
			}
		} else {
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

func isSafe(sli *[]int) bool {
	isDesc := sort.SliceIsSorted(*sli, func(x, y int) bool {
		return (*sli)[x] > (*sli)[y]
	})
	isAsc := sort.IntsAreSorted(*sli)

	if isAsc || isDesc {
		for i := 1; i < len(*sli); i++ {
			distance := abs((*sli)[i-1], (*sli)[i])
			if distance < 1 || distance > 3 {
				break
			}
			if i == len(*sli)-1 {
				return true
			}
		}
	}
	return false
}
