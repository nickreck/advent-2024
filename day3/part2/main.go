package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	content, _ := os.ReadFile("../input.txt")
	str := string(content)

	re := regexp.MustCompile("((mul\\()[0-9]+,[0-9]+\\))|don't\\(\\)|do\\(\\)")
	matches := re.FindAllString(str, -1)

	re = regexp.MustCompile("[0-9]+")
	sum := 0
	active := true
	for _, v := range matches {
		if v == "don't()" {
			active = false
			continue
		} else if v == "do()" {
			active = true
			continue
		}

		if !active {
			continue
		}

		sli := re.FindAllString(v, -1)
		x, _ := strconv.Atoi(sli[0])
		y, _ := strconv.Atoi(sli[1])
		sum += x * y
	}
	fmt.Println(sum)
}
