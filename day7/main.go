package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	problems := strings.Split(strings.TrimSuffix(string(input), "\n"), "\n")
	m := make(map[int][][]int)
	for _, v := range problems {
		sli := strings.Split(v, ": ")
		key, _ := strconv.Atoi(sli[0])
		sli = strings.Fields(sli[1])
		integers := make([]int, len(sli))
		for i, v := range sli {
			n, _ := strconv.Atoi(v)
			integers[i] = n
		}
		if v, ok := m[key]; ok {
			v = append(v, integers)
			m[key] = v
		} else {
			m[key] = [][]int{integers}
		}
	}

	sum := 0
	for k, v := range m {
		for _, integers := range v {
			for i := 0; i < (1 << (len(integers) - 1)); i++ {
				str := fmt.Sprintf("%0*b", len(integers)-1, i)
				total := integers[0]
				for ci, c := range str {
					if c == '0' {
						total *= integers[ci+1]
					} else {
						total += integers[ci+1]
					}

					if total > k {
						break
					}
				}
				if total == k {
					sum += total
					break
				}
			}
		}
	}
	fmt.Println(sum)
}
