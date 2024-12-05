package main

import (
	"bufio"
	"fmt"
	"os"
)

func part2() {
	file, _ := os.Open("input.txt")
	reader := bufio.NewReader(file)
	puzzle := make([][]rune, 0)
	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			break
		}
		puzzle = append(puzzle, []rune(string(line)))
	}

	sum := 0
	for c := 1; c < len(puzzle)-1; c++ {
		for r := 1; r < len(puzzle[0])-1; r++ {
			if puzzle[c][r] == 'A' {
				if puzzle[c-1][r-1] == 'M' && puzzle[c+1][r+1] == 'S' && puzzle[c+1][r-1] == 'M' && puzzle[c-1][r+1] == 'S' {
					sum++
				}
				if puzzle[c-1][r-1] == 'S' && puzzle[c+1][r+1] == 'M' && puzzle[c+1][r-1] == 'S' && puzzle[c-1][r+1] == 'M' {
					sum++
				}
				if puzzle[c-1][r-1] == 'M' && puzzle[c+1][r-1] == 'S' && puzzle[c-1][r+1] == 'M' && puzzle[c+1][r+1] == 'S' {
					sum++
				}
				if puzzle[c-1][r-1] == 'S' && puzzle[c+1][r-1] == 'M' && puzzle[c-1][r+1] == 'S' && puzzle[c+1][r+1] == 'M' {
					sum++
				}
			}
		}
	}
	fmt.Println(sum)
}
