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

	re := regexp.MustCompile(`(?s)don't\(\).*?do\(\)`)
	str = re.ReplaceAllString(str, "-")

	re = regexp.MustCompile(`((mul\()\d+,\d+\))`)
  matches := re.FindAllString(str, -1)

	re = regexp.MustCompile(`\d+`)
	sum := 0
	for _, v := range matches {
		sli := re.FindAllString(v, -1)
		x, _ := strconv.Atoi(sli[0])
		y, _ := strconv.Atoi(sli[1])
		sum += x * y
	}
	fmt.Println(sum)
}
