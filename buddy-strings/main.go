package main

import "fmt"

func main() {
	fmt.Println(buddyStrings("abba", "aabb"))
}

// LeetCode 859
func buddyStrings(s string, goal string) bool {
	if len(s) == 1 || (len(s) == 2 && s[0] == goal[0] && s[1] == goal[1] && s[0] != goal[1]) || len(s) != len(goal) {
		return false
	}

	n := len(s)
	repeatedLetter := false

	if n > 2 && s == goal {
		runeMap := make(map[rune]int)
		for _, v := range s {
			_, ok := runeMap[v]
			if ok {
				repeatedLetter = true
			}
			runeMap[v]++
		}
		if !repeatedLetter {
			return false
		}
	}

	var sTemp byte
	var gTemp byte
	var count int

	for i := 0; i < n; i++ {
		if s[i] != goal[i] {
			count++
			if sTemp == 0 && gTemp == 0 {
				sTemp = s[i]
				gTemp = goal[i]
			} else if count > 2 || sTemp != goal[i] || gTemp != s[i] {
				return false
			}
		}
	}
	if count == 1 && sTemp != gTemp {
		return false
	}
	return true
}
