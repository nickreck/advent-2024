package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	n := flag.Int("part1", 25, "Number of blinks for part 1")
	n2 := flag.Int("part2", 75, "Number of blinks for part 1")
	flag.Parse()

	rocks := loadData()
	m := make(map[node]int)
  start := time.Now().Nanosecond()
	blink(rocks, *n, m)
  end := time.Now().Nanosecond()
  fmt.Printf("Part 1: %d\n", end - start)
  start = time.Now().Nanosecond()
	blink(rocks, *n2, m)
  end = time.Now().Nanosecond()
  fmt.Printf("Part 2: %d\n", end - start)
}

func blink(rocks []int, n int, m map[node]int) {
	sum := 0
	for _, v := range rocks {
		sum += applyRule(v, n, m)
	}
	fmt.Println(sum)
}

func applyRule(v int, n int, m map[node]int) int {
	if count, exists := m[node{v, n}]; exists {
		return count
	}
	var ret int
	if n == 0 {
		ret = 1
	} else if v == 0 {
		ret = applyRule(1, n-1, m)
	} else if intLen(v)%2 == 0 {
		s := strconv.Itoa(v)
		r1, _ := strconv.Atoi(s[:len(s)/2])
		r2, _ := strconv.Atoi(s[len(s)/2:])
		ret = applyRule(r1, n-1, m) + applyRule(r2, n-1, m)
	} else {
		ret = applyRule(v * 2024, n-1, m)
	}
	m[node{v, n}] = ret
	return ret
}

type node struct {
	v int
	c int
}

func loadData() []int {
	file, _ := os.ReadFile("input.txt")
	str := strings.Fields(strings.TrimSuffix(string(file), "\n"))
	var data []int
	for _, v := range str {
		n, _ := strconv.Atoi(v)
		data = append(data, n)
	}
	return data
}

func intLen(n int) int {
	count := 0
	for n > 0 {
		n = n / 10
		count++
	}
	return count
}
