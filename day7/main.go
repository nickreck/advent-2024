package main

import (
	"fmt"
	"math"
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

	base2, base3 := findSum(m)

	fmt.Println(base2)
	fmt.Println(base3)
}

func findSum(m map[int][][]int) (int, int) {
	sum := 0
	base3Sum := 0
	for k, v := range m {
		for _, integers := range v {
			bitwise(k, &sum, integers)
			base3(k, &base3Sum, integers)
		}
	}
	return sum, base3Sum
}

func base3(k int, sum *int, integers []int) {
	for i := 0; i < (int(math.Pow(float64(3), float64(len(integers)-1)))); i++ {
    numStr := strconv.FormatInt(int64(i), 3)
    str := fmt.Sprintf("%*s", len(integers)-1, numStr)
		total := integers[0]
		for ci, c := range str {
			if c == '0' || c == ' ' {
				total *= integers[ci+1]
			} else if c == '1' {
				total += integers[ci+1]
			} else {
        total, _ = strconv.Atoi(fmt.Sprintf("%d%d", total, integers[ci+1]))
      }

			if total > k {
				break
			}
		}
		if total == k {
			*sum += total
			break
		}
	}
}

func bitwise(k int, sum *int, integers []int) {
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
			*sum += total
			break
		}
	}
}
