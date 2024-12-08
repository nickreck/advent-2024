package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	inputRaw, _ := os.ReadFile("input.txt")
	input := strings.Split(strings.TrimSuffix(string(inputRaw), "\n"), "\n")
	m := make(map[rune][]coordinates)
	nodes := make(map[coordinates]struct{})
	count := 0
	for y, s := range input {
		for x, c := range s {
			if c != '.' && c != '#' {
				if arr, exists := m[c]; exists {
					coords := coordinates{x, y}
					findAntinodes(arr, coords, nodes, &count, input)
					arr = append(arr, coords)
					m[c] = arr
				} else {
					m[c] = []coordinates{{x, y}}
				}
			}
		}
	}
	fmt.Println(count)
}

func findAntinodes(arr []coordinates, current coordinates, nodes map[coordinates]struct{}, int *int, input []string) {
	for _, coord := range arr {
		if _, exists := nodes[current]; !exists {
			(*int)++
			nodes[current] = struct{}{}
		}
		if _, exists := nodes[coord]; !exists {
			(*int)++
			nodes[coord] = struct{}{}
		}
		copyCurr := current
		dx := copyCurr.x - coord.x
		dy := copyCurr.y - coord.y

		valid := true
		for valid {
			valid, copyCurr = validateAntinode(copyCurr, nodes, int, dx, dy, input)
		}

		dx *= -1
		dy *= -1
		valid = true
		for valid {
			valid, coord = validateAntinode(coord, nodes, int, dx, dy, input)
		}
	}
}

func validateAntinode(coord coordinates, nodes map[coordinates]struct{}, int *int, dx, dy int, input []string) (bool, coordinates) {
	node := coordinates{coord.x + dx, coord.y + dy}
	in := isNodeInBounds(node, input)
	if _, exists := nodes[node]; !exists && in {
		(*int)++
		nodes[node] = struct{}{}
	}
	return in, node
}

func isNodeInBounds(node coordinates, input []string) bool {
	return node.x >= 0 && node.x < len(input[0]) && node.y >= 0 && node.y < len(input)
}

type coordinates struct {
	x int
	y int
}
