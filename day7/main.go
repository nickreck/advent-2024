package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	input, _ := os.ReadFile("input.txt")
	problems := strings.Split(strings.TrimSuffix(string(input), "\n"), "\n")
	m := make(map[int][][]int)
	for _, v := range problems {
		sli := strings.Split(v, ": ")
		key, _ := strconv.Atoi(sli[0])
		sli = strings.Fields(sli[1])
		integers := make([]int, len(sli))
		for i, v := range sli {
			n, _ := strconv.Atoi(v)
			integers[i] = n
		}
		if v, ok := m[key]; ok {
			v = append(v, integers)
			m[key] = v
		} else {
			m[key] = [][]int{integers}
		}
	}

	base2, base3 := findSum(m)

	fmt.Println(base2)
	fmt.Println(base3)
}

func findSum(m map[int][][]int) (int, int) {
	base2Chan := make(chan int)
	base3Chan := make(chan int)
	wg1 := &sync.WaitGroup{}
	wg2 := &sync.WaitGroup{}
	for k, v := range m {
		for _, integers := range v {
			wg1.Add(1)
			go bitwise(wg1, k, base2Chan, integers)
			wg2.Add(1)
			go base3(wg2, k, base3Chan, integers)
		}
	}

	go func() {
		defer close(base2Chan)
		wg1.Wait()
	}()
	go func() {
		defer close(base3Chan)
		wg2.Wait()
	}()

	base2Sum := 0
	base3Sum := 0
	for i := range base3Chan {
		base3Sum += i
	}
	for i := range base2Chan {
		base2Sum += i
	}
	return base2Sum, base3Sum
}

func base3(wg *sync.WaitGroup, k int, c chan int, integers []int) {
	defer wg.Done()
	for i := 0; i < (int(math.Pow(float64(3), float64(len(integers)-1)))); i++ {
		numStr := strconv.FormatInt(int64(i), 3)
		str := fmt.Sprintf("%*s", len(integers)-1, numStr)
		total := checkRules(k, integers, str)
		if total == integers[0] {
			c <- k
			break
		}
	}
}

func bitwise(wg *sync.WaitGroup, k int, c chan int, integers []int) {
	defer wg.Done()
	for i := 0; i < (1 << (len(integers) - 1)); i++ {
		str := fmt.Sprintf("%0*b", len(integers)-1, i)
		total := checkRules(k, integers, str)
		if total == integers[0] {
			c <- k
			break
		}
	}
}

func checkRules(k int, integers []int, str string) int {
	total := k
layerFor:
	for j := len(str) - 1; j >= 0; j-- {
		switch str[j] {
		case '0', ' ':
			if total%integers[j+1] != 0 {
				break layerFor
			}
			total /= integers[j+1]
			break
		case '1':
			total -= integers[j+1]
			break
		case '2':
			hold := strconv.Itoa(total)
			hL := len(hold)
			iL := intLen(integers[j+1])
			if hL < iL {
				break layerFor
			}
			hold = hold[hL-iL:]
			check, _ := strconv.Atoi(hold)
			if check != integers[j+1] {
				break layerFor
			}
			hold = strconv.Itoa(total)
			hold = hold[:len(hold)-intLen(integers[j+1])]
			total, _ = strconv.Atoi(hold)
			break
		}

		if total < 0 {
			break
		}
	}
	return total
}

func intLen(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
}
