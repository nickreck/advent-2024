package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func part1() {
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
		for i := len(sli) - 1; i >= 0 && valid; i-- {
			val, _ := strconv.Atoi(sli[i])
			constraints := rules[val]
			for j := i - 1; j >= 0 && valid; j-- {
				check, _ := strconv.Atoi(sli[j])
				if contains(constraints, check) {
					valid = false
				}
			}
		}

		if !valid {
			continue
		}

		index := len(sli) / 2
		add, _ := strconv.Atoi(sli[index])
		sum += add
	}

	fmt.Println(sum)
}

func contains(sli []int, val int) bool {
	for _, v := range sli {
		if v == val {
			return true
		}
	}
	return false
}
