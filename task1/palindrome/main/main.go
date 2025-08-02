package main

import "fmt"

// 判断一个整数是否是回文数
func isPalindrome(x int) bool {
	// 负数，正数能被0整除都不是回文数
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	// 将数字进行反转，且只反转一半
	reverted := 0
	for x > reverted {
		reverted = reverted*10 + x%10
		x /= 10
	}

	//数字个数，奇数位（将反转数除以10）｜ 偶数位处理(直接比对)
	return x == reverted || x == reverted/10
}

func main() {
	// 回文数
	palindrome := isPalindrome(122)
	fmt.Printf("palindrome: %+v\n", palindrome)
	b := isPalindrome(1221)
	fmt.Printf("b: %+v\n", b)
	b2 := isPalindrome(0)
	fmt.Printf("b2: %+v\n", b2)
}
