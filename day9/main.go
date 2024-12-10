package main

import (
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

func main() {
	inputRaw, _ := os.ReadFile("input.txt")
	input := []rune(strings.TrimSuffix(string(inputRaw), "\n"))

	s := make(map[int]int)
	indexToAddAt := 0
	id := 0
	disk := make([]int, 0)
	sum := 0
	for i := 0; i < len(input); i += 2 {
		count, _ := strconv.Atoi(string(input[i]))
		for j := 0; j < count; j++ {
			s[id]++
			disk = append(disk, id)
			sum += indexToAddAt * id
			indexToAddAt++
		}
		id++

		if i+1 < len(input) {
			count, _ = strconv.Atoi(string(input[i+1]))
			for j := 0; j < count; j++ {
				disk = append(disk, -1)
			}
		}
	}

  // part1(disk, s)
  part2(disk, s)


	// lastId := ((len(input) + 1) / 2) - 1
	// countOfLast, _ := strconv.Atoi(string(input[len(input)-1]))
	// indexToAddAt := 0
	// id := 0
	// rev := len(input) - 1
	// disk := make([]int, 0)
	// sum := 0
	// for i := 0; i < rev; i += 2 {
	// 	count, _ := strconv.Atoi(string(input[i]))
	// 	for j := 0; j < count; j++ {
	// 		disk = append(disk, id)
	// 		sum += indexToAddAt * id
	// 		indexToAddAt++
	// 	}
	// 	id++

	// 	// While there's free space
	// 	count, _ = strconv.Atoi(string(input[i+1]))
	// 	for count > 0 && rev > i {
	// 		disk = append(disk, lastId)
	// 		sum += indexToAddAt * lastId
	// 		indexToAddAt++
	// 		countOfLast--
	// 		count--
	// 		if countOfLast == 0 {
	// 			lastId--
	// 			rev -= 2
	// 			countOfLast, _ = strconv.Atoi(string(input[rev]))
	// 		}
	// 	}
	// }
	// for countOfLast > 0 {
	// 	disk = append(disk, lastId)
	// 	sum += indexToAddAt * lastId
	// 	indexToAddAt++
	// 	countOfLast--
	// }
	check := 0
	for i, v := range disk {
		if v != -1 {
			check += i * v
		}
	}
	// fmt.Println(sum)
	fmt.Println(check)
	// fmt.Println(disk)
}

func part2(disk []int, s map[int]int) {
  sli := make([]int, 0)
  for k := range s {
    sli = append(sli, k)
  }
  sort.Sort(sort.Reverse(sort.IntSlice(sli)))
  
  for _, v := range sli {
    size := s[v]
    end := slices.Index(disk, v)
    for i := 0; i <= end-size; i++ {
      index := i
      count := 0
      for disk[index] == -1 && count < size{
        count++
        index++
      }
      if count == size {
        for j := slices.Index(disk, v); j != -1; j = slices.Index(disk, v) {
          disk[j] = -1
        }
        for count = 0; count < size; count++ {
          disk[i] = v
          i++
        }
        break
      }
    }
  }
}

func part1(disk []int, s map[int]int) {
  lastIndex := len(disk) - 1
  for i, v := range disk {
    if s[v] > 0 {
      s[v]--
    }
    if v == -1 && len(s) > 0 {
      for disk[lastIndex] == -1 {
        lastIndex--
      }
      s[disk[lastIndex]]--
      if s[disk[lastIndex]] == 0 {
        delete(s, disk[lastIndex])
      }
      disk[i] = disk[lastIndex]
      disk[lastIndex] = -1
      lastIndex--
    }
    if s[v] == 0 {
      delete(s, v)
    }
  }
}
