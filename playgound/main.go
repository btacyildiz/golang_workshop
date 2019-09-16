package main

import "strconv"

func Valid(num string) bool {
	numSum := 0
	numCount := 0

	for i := len(num) - 1; i >= 0; i-- {
		ch := num[i]
		println(string(ch))
		if ch == ' ' {
			continue
		}
		num, err := strconv.Atoi(string(ch))
		if err != nil {
			// reject other than space characters
			println("rejected: ", err)
			return false
		}
		numCount++
		if numCount%2 == 0 {
			if num*2 > 9 {
				numSum += num*2 - 9
			} else {
				numSum += num * 2
			}
		} else {
			numSum += num
		}
	}
	if numSum%10 != 0 {
		return false
	}
	return true
}

func main() {
	println("Is Valid: ", Valid("059"))
}
