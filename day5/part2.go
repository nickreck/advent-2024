package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func part2() {
	file, _ := os.Open("input.txt")
	r := bufio.NewReader(file)

	rules := make(map[int][]int)
	for {
		orderingRulesRaw, err := r.ReadString('\n')
		orderingRulesRaw = strings.TrimSuffix(orderingRulesRaw, "\n")
		if err != nil || orderingRulesRaw == "" {
			break
		}

		sli := strings.Split(orderingRulesRaw, "|")
		before, _ := strconv.Atoi(sli[0])
		after, _ := strconv.Atoi(sli[1])
		beforeList := rules[before]
		if beforeList == nil {
			rules[before] = []int{after}
		} else {
			rules[before] = append(beforeList, after)
		}
	}

	sum := 0
	for {
		listForValidation, err := r.ReadString('\n')
		listForValidation = strings.TrimSuffix(listForValidation, "\n")
		if err != nil || listForValidation == "" {
			break
		}

		sli := strings.Split(listForValidation, ",")
		valid := true
		for i := len(sli) - 1; i > 0 && valid; i-- {
			val, _ := strconv.Atoi(sli[i])
			constraints := rules[val]
			for j := i - 1; j >= 0 && valid; j-- {
				check, _ := strconv.Atoi(sli[j])
				if slices.Index(constraints, check) != -1 {
					valid = false
				}
			}
		}

		if valid {
			continue
		}

		ordered := fixOrdering(sli, &rules)

		index := len(ordered) / 2
		sum += ordered[index]
	}

	fmt.Println(sum)
}

func fixOrdering(sli []string, m *map[int][]int) []int {
	ordered := make([]int, 0)
	for i := 0; i < len(sli); i++ {
		min := len(ordered)
		val, _ := strconv.Atoi(sli[i])
		constraints := (*m)[val]
		for _, c := range constraints {
			index := slices.Index(ordered, c)
			if index != -1 && min > index {
				min = index
			}
		}
		if min == len(ordered) {
			ordered = append(ordered, val)
		} else {
			ordered = slices.Insert(ordered, min, val)
		}
	}
	return ordered
}

func tryAndPlace(int *[]int, val, index int) {
	if (*int)[index] == 0 {
		(*int)[index] = val
		return
	}
	index--
	tryAndPlace(int, val, index)
}
