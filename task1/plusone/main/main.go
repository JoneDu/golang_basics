package main

import "fmt"

func main() {
	one := plusOne([]int{3, 4, 5, 9})
	fmt.Printf("one: %+v\n", one)
}

func plusOne(digits []int) []int {
	//这个问题，主要是加一，和进位的问题。从最低位开始处理
	l := len(digits)
	for i := (l - 1); i >= 0; i-- {
		digits[i]++
		//检查是否进位
		if digits[i] < 10 {
			return digits
		}
		// 当前位为0，继续处理前一位
		digits[i] = 0
	}

	return append([]int{1}, digits...)
}
