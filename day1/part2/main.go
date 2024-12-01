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
	defer file.Close()

	reader := bufio.NewReader(file)

	similarity := 0
	leftSli := make([]int, 0)
	rightSli := make([]int, 0)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		sli := strings.Fields(line)
		left, _ := strconv.Atoi(sli[0])
		leftSli = append(leftSli, left)
		right, _ := strconv.Atoi(sli[1])
		rightSli = append(rightSli, right)
	}

  sort.Ints(leftSli)
  sort.Ints(rightSli)

	for _, v := range leftSli {
		index := sort.SearchInts(rightSli, v)
    count := 0
    for index < len(rightSli) && rightSli[index] == v {
      count++
      index++
    }
    similarity += v * count
	}

	fmt.Println(similarity)
}
