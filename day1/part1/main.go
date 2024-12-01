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

	distance := 0
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

	for i := 0; i < len(leftSli); i++ {
		distance += absDiffInt(leftSli[i], rightSli[i])
	}
	fmt.Println(distance)
}
func absDiffInt(x, y int) int {
   if x < y {
      return y - x
   }
   return x - y
}
