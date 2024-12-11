package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	input := loadInput()
	start := time.Now().UnixNano()
	bothParts(input)
	end := time.Now().UnixNano()
	fmt.Println(end - start)
}

func bothParts(input *[][]rune) {
	data := *input
	sum := 0
	sum2 := 0
	for r, rv := range data {
		for c, cv := range rv {
			if cv == '0' {
				m := make(map[coordinate]struct{})
				num, _ := strconv.Atoi(string(cv))
				sum += hike(r, c, num, data, m, false)
				sum2 += hike(r, c, num, data, m, true)
			}
		}
	}
	fmt.Println(sum)
  fmt.Println(sum2)
}

func hike(r, c, num int, data [][]rune, m map[coordinate]struct{}, part2 bool) int {
	sum := 0
	check, _ := strconv.Atoi(string(data[r][c]))
	if check != num {
		return 0
	} else if check == 9 {
		if _, exists := m[coordinate{r, c}]; !exists || part2 {
			m[coordinate{r, c}] = struct{}{}
			return 1
		}
		return 0
	}

	if r > 0 {
		sum += hike(r-1, c, num+1, data, m, part2)
	}
	if r < len(data)-1 {
		sum += hike(r+1, c, num+1, data, m, part2)
	}
	if c > 0 {
		sum += hike(r, c-1, num+1, data, m, part2)
	}
	if c < len(data[0])-1 {
		sum += hike(r, c+1, num+1, data, m, part2)
	}

	return sum
}

type coordinate struct {
	r int
	c int
}

func loadInput() *[][]rune {
	file, _ := os.ReadFile("input.txt")
	inputRaw := strings.Split(strings.TrimSuffix(string(file), "\n"), "\n")
	input := make([][]rune, 0)
	for _, s := range inputRaw {
		input = append(input, []rune(s))
	}
	return &input
}
