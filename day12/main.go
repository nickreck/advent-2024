package main

import (
	"fmt"
	"os"
	"strings"
)

var history = make(map[coordinate]struct{})

func main() {
	file, _ := os.ReadFile("input.txt")
	input := strings.Split(strings.TrimSuffix(string(file), "\n"), "\n")

	var sum int
	for r := range input {
		for c := range input[r] {
			currentLoc := coordinate{c, r}
			if _, exists := history[currentLoc]; !exists {
				area := 0
				count := searchPlot(input, currentLoc, &area)
				sum += area * count
			}
		}
	}
	fmt.Println(sum)
}

func searchPlot(input []string, c coordinate, area *int) int {
	if _, exists := history[c]; exists {
		return 0
	} else {
		history[c] = struct{}{}
		(*area)++
	}

	sum := 0
	char := []rune(input[c.y])[c.x]

	// Check Right
	newLoc := coordinate{c.x + 1, c.y}
	sum += getPerimeter(newLoc, input, area, char)

	// Check Down
	newLoc = coordinate{c.x, c.y + 1}
	sum += getPerimeter(newLoc, input, area, char)

	// Check Up
	newLoc = coordinate{c.x, c.y - 1}
	sum += getPerimeter(newLoc, input, area, char)

	// Check Left
	newLoc = coordinate{c.x - 1, c.y}
	sum += getPerimeter(newLoc, input, area, char)

	return sum
}

func getPerimeter(c coordinate, input []string, area *int, char rune) int {
	sum := 0
	valid := getValidNeighbor(c, input)
	if char == valid {
		sum += searchPlot(input, c, area)
	} else {
		sum++
	}
	return sum
}

func getValidNeighbor(c coordinate, input []string) rune {
	if c.x < len(input[0]) && c.x >= 0 && c.y < len(input) && c.y >= 0 {
		return []rune(input[c.y])[c.x]
	}
	return ' '
}

type coordinate struct {
	x int
	y int
}
