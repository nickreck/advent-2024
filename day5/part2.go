package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part2() {
  input, _ := os.ReadFile("input.txt")
  split := strings.Split(strings.TrimSpace(string(input)), "\n\n")

  comparator := func(x, y string) int {
    for _, v := range strings.Split(split[0], "\n") {
      rule := strings.Split(v, "|")
      if rule[0] == x && rule[1] == y {
        return -1
      }
    }
    return 0
  }

  count := 0
  for _, v := range strings.Split(split[1], "\n") {
    sli := strings.Split(v, ",")
    if !slices.IsSortedFunc(sli, comparator) {
      slices.SortFunc(sli, comparator)
      add, _ := strconv.Atoi(sli[len(sli)/2])
      count += add
    }
  }
  fmt.Println(count)
}
