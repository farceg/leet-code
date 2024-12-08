package moneydistribution

import "fmt"

func DistMoney(money int, children int) int {
	// Cada niño recibe 1 dolar
	money -= children
	if money < 7 {
		if money < 0 {
			return -1
		}
		return 0
	}
	maxiWithEight, remainings := money/7, money%7
	fmt.Println(maxiWithEight, remainings)

	if maxiWithEight >= children {
		if maxiWithEight == children && remainings == 0 {
			return maxiWithEight
		}
		return children - 1
	} else if children == maxiWithEight+1 && remainings == 3 {
		return maxiWithEight - 1
	}

	return maxiWithEight
}
