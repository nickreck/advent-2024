package main

import (
	"bufio"
	"fmt"
	"os"
)

func part1() {
	file, _ := os.Open("../input.txt")
	reader := bufio.NewReader(file)
	puzzle := make([]string, 0)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		puzzle = append(puzzle, string(line))
	}

	sum := 0
	for c := 0; c < len(puzzle); c++ {
		runeArr := []rune(puzzle[c])
		for r := 0; r < len(runeArr); r++ {
			if runeArr[r] == 'X' {
				if r > 2 {
					sum += findChristmas(puzzle, r-1, c, -1, 0, 1)
				}
				if r < len(puzzle[0])-4 {
					sum += findChristmas(puzzle, r+1, c, 1, 0, 1)
				}
				if c < len(puzzle)-4 {
					sum += findChristmas(puzzle, r, c+1, 0, 1, 1)
				}
				if c > 2 {
					sum += findChristmas(puzzle, r, c-1, 0, -1, 1)
				}
				if c > 2 && r > 2 {
					sum += findChristmas(puzzle, r-1, c-1, -1, -1, 1)
				}
				if c < len(puzzle)-4 && r > 2 {
					sum += findChristmas(puzzle, r-1, c+1, -1, 1, 1)
				}
				if c < len(puzzle)-4 && r < len(puzzle[0])-4 {
					sum += findChristmas(puzzle, r+1, c+1, 1, 1, 1)
				}
				if c > 2 && r < len(puzzle[0])-4 {
					sum += findChristmas(puzzle, r+1, c-1, 1, -1, 1)
				}
			}
		}
	}
	fmt.Println(sum)
}

func findChristmas(puzzle []string, r, c, dr, dc, index int) int {
	key := []rune{'X', 'M', 'A', 'S'}
	char := []rune(puzzle[c])
	if char[r] != key[index] {
		return 0
	} else if index == 3 {
		return 1
	}
	index++
	return findChristmas(puzzle, r+dr, c+dc, dr, dc, index)
}
