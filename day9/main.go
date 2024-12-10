package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	inputRaw, _ := os.ReadFile("input.txt")
	input := []rune(strings.TrimSuffix(string(inputRaw), "\n"))

	start := time.Now().UnixNano()
	lastId := ((len(input) + 1) / 2) - 1
	countOfLast, _ := strconv.Atoi(string(input[len(input)-1]))
	indexToAddAt := 0
	id := 0
	rev := len(input) - 1
	disk := make([]int, 0)
	sum := 0
	for i := 0; i < rev; i += 2 {
		count, _ := strconv.Atoi(string(input[i]))
		for j := 0; j < count; j++ {
			disk = append(disk, id)
			sum += indexToAddAt * id
			indexToAddAt++
		}
		id++

		// While there's free space
		count, _ = strconv.Atoi(string(input[i+1]))
		for count > 0 && rev > i {
			disk = append(disk, lastId)
			sum += indexToAddAt * lastId
			indexToAddAt++
			countOfLast--
			count--
			if countOfLast == 0 {
				lastId--
				rev -= 2
				if rev <= i {
					break
				}
				countOfLast, _ = strconv.Atoi(string(input[rev]))
			}
		}
	}
	for countOfLast > 0 {
		disk = append(disk, lastId)
		sum += indexToAddAt * lastId
		indexToAddAt++
		countOfLast--
	}
	end := time.Now().UnixNano()
	fmt.Println(sum)
	fmt.Println(end - start)
}
