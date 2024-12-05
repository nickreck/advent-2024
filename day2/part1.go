package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func part1() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewReader(file)

	count := 0
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

		isDesc := sort.SliceIsSorted(intSli, func(x, y int) bool {
			return intSli[x] > intSli[y]
		})
		isAsc := sort.IntsAreSorted(intSli)

		if isAsc || isDesc {
			for i := 1; i < len(intSli); i++ {
				distance := abs(intSli[i-1], intSli[i])
				if distance < 1 || distance > 3 {
					break
				}
				if i == len(intSli)-1 {
					count++
				}
			}
		}
	}
	fmt.Println(count)
}

func abs(x, y int) int {
	if x > y {
		return x - y
	}
	return y - x
}
