package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()

	reader := bufio.NewReader(file)

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

	m := make(map[int]int)
	for _, value := range rightSli {
		m[value]++
	}

	similarity := 0
	for _, value := range leftSli {
		similarity += value * m[value]
	}

	fmt.Println(similarity)
}
