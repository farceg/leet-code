package main

import "fmt"

func main() {
	s := "aabaaaacaabc"
	k := 2
	fmt.Println("Return:", takeCharacters(s, k))

	s = "a"
	k = 1
	fmt.Println("Return:", takeCharacters(s, k))

	s = "a"
	k = 0
	fmt.Println("Return:", takeCharacters(s, k))

	s = "abc"
	k = 1
	fmt.Println("Return:", takeCharacters(s, k))
}

func takeCharacters(s string, k int) int {

	if k == 0 {
		return 0
	}
	if len(s) < 3 {
		return -1
	}

	cMap := map[rune]int{'a': 0, 'b': 0, 'c': 0}
	for _, v := range s {
		cMap[v]++
	}

	if cMap['a'] < k || cMap['b'] < k || cMap['c'] < k {
		return -1
	}

	res := 100001
	left := 0

	for i, key := range s {
		cMap[key]--
		for min(cMap) < k {
			cMap[key]++
			left++
		}
		if res > len(s)-(i-left+1) {
			res = len(s) - (i - left + 1)
		}
	}

	return res

}

func min(cMap map[rune]int) int {
	min := 100001
	for _, v := range cMap {
		if v < min {
			min = v
		}
	}
	return min
}
