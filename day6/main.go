package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fileContents, _ := os.ReadFile("input.txt")
	guardPos := strings.Index(string(fileContents), "^")
	pathMap := strings.Split(strings.TrimSuffix(string(fileContents), "\n"), "\n")
	row := int((float32(guardPos) / float32(len(fileContents))) * float32(len(pathMap)))
	col := guardPos % (len(pathMap[0]) + 1)
	startingPoint := point{col, row, '^'}

	count, m := traverseFromStart(startingPoint, pathMap)
	trapCount := trapHim(startingPoint, m, pathMap)
	fmt.Println(count)
	fmt.Println(trapCount)
}

type point struct {
	x         int
	y         int
	direction byte
}

type coordinates struct {
	x int
	y int
}

func traverseFromStart(loc point, pathMap []string) (int, map[coordinates]struct{}) {
	pathTraveled := make(map[coordinates]struct{})
	count := 0
	for {
		if _, exists := pathTraveled[coordinates{loc.x, loc.y}]; !exists {
			count++
			pathTraveled[coordinates{loc.x, loc.y}] = struct{}{}
		}

		ok, newLoc := attemptMove(loc, pathMap)
		if !ok {
			return count, pathTraveled
		}
		loc = newLoc
	}
}

func attemptMove(originalLoc point, pathMap []string) (bool, point) {
	newLoc := originalLoc
	switch newLoc.direction {
	case '^':
		newLoc.y -= 1
	case '>':
		newLoc.x += 1
	case 'v':
		newLoc.y += 1
	case '<':
		newLoc.x -= 1
	}

	if newLoc.x < 0 || newLoc.y < 0 || newLoc.x >= len(pathMap[0]) || newLoc.y >= len(pathMap) {
		return false, newLoc
	}

	switch pathMap[newLoc.y][newLoc.x] {
	case '#':
		switch originalLoc.direction {
		case '^':
			originalLoc.direction = '>'
		case '>':
			originalLoc.direction = 'v'
		case 'v':
			originalLoc.direction = '<'
		case '<':
			originalLoc.direction = '^'
		}
		return attemptMove(originalLoc, pathMap)
	case '.':
		return true, newLoc
	case '^':
		return true, newLoc
	}
	return false, newLoc
}

func trapHim(loc point, m map[coordinates]struct{}, pathMap []string) int {
  count := 0
	for k, _ := range m {
		str := pathMap[k.y]
    hold := str
		str = str[:k.x] + "#" + str[k.x+1:]
    pathMap[k.y] = str
		count += traverseInfinitely(loc, pathMap)
    pathMap[k.y] = hold
	}
  return count
}

func traverseInfinitely(loc point, pathMap []string) int {
	pathTraveled := make(map[point]struct{})
	count := 0
	for {
		if _, exists := pathTraveled[loc]; !exists {
			count++
			pathTraveled[loc] = struct{}{}
		} else {
      return 1
    }

		ok, newLoc := attemptMove(loc, pathMap)
		if !ok {
			return 0
		}
		loc = newLoc
	}
}
