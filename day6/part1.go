package main

import (
	"fmt"
	"os"
	"slices"
	"strings"
)

func part1() {
	fileContents, _ := os.ReadFile("input.txt")
	guardPos := strings.Index(string(fileContents), "^")
	pathMap := strings.Split(strings.TrimSuffix(string(fileContents), "\n"), "\n")
	row := int((float32(guardPos) / float32(len(fileContents))) * float32(len(pathMap)))
	col := guardPos % (len(pathMap[0]) + 1)

	count := 1
	tryMoving(row, col, &count, &pathMap)

	fmt.Println(count)
}

func putAnX(row, col int, arr *[]string) {
	str := []byte((*arr)[row])
	str[col] = 'X'
	(*arr)[row] = string(str)
}

func tryMoving(row, col int, int *int, arr *[]string) {
	cur := (*arr)[row][col]
	r, c, safe := isSafeMove(cur, row, col, arr)
	if !safe && inArea(r, c, arr) {
		directions := []byte{'^', '>', 'v', '<'}
		index := slices.Index(directions, cur)
		directions = append(directions[index:], directions[:index]...)
		for _, v := range directions[1:] {
			if tr, tc, safe := isSafeMove(v, row, col, arr); safe {
				r = tr
				c = tc
				cur = v
				break
			}
			return
		}
	} else if !safe {
		return
	}
  if (*arr)[r][c] != 'X' {
    (*int)++
  }
	putAnX(row, col, arr)
	putAChar(r, c, arr, cur)
	tryMoving(r, c, int, arr)
}

func putAChar(row, col int, arr *[]string, x byte) {
	str := []byte((*arr)[row])
	str[col] = x
	(*arr)[row] = string(str)
}

func isSafeMove(char byte, r, c int, arr *[]string) (int, int, bool) {
	if char == '^' {
		r--
	} else if char == '>' {
		c++
	} else if char == '<' {
		c--
	} else {
		r++
	}

	return r, c, inArea(r, c, arr) && (*arr)[r][c] != '#'
}

func inArea(r, c int, arr *[]string) bool {
	return r >= 0 && r < len(*arr) && c >= 0 && c < len(*arr)
}
