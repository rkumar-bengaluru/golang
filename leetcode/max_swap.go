package leetcode

import (
	"fmt"
	"strconv"
)

func maximumSwap(n int) (res int) {
	s := fmt.Sprint(n)
	chars := []rune(s)
	char_map := make(map[rune]int)
	for i, v := range chars {
		char_map[v] = i
	}
	digits := [10]rune{'9', '8', '7', '6', '5', '4', '3', '2', '1', '0'}

	for i, v := range chars {

		fmt.Println(i, v)
		for _, digit := range digits {
			if digit <= v {
				break
			}
			val, foundInMap := char_map[rune(digit)]
			if foundInMap && val > i {
				chars[i], chars[char_map[rune(digit)]] = rune(digit), chars[i]
				tmp := ""
				for _, c := range chars {
					tmp += string(c)
				}
				res, _ := strconv.Atoi(tmp)
				return res
			}
		}
	}
	return n
}
