package main

import "fmt"

func main() {
	fmt.Println(isValid("("))
	fmt.Println(isValid("()"))
	fmt.Println(isValid("()[]"))
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("({[]})"))
	fmt.Println(isValid("({[}])"))
}

func isValid(s string) bool {
	// 使用栈的思路，将左边的括号，压入栈中，右边的括号进行和左边的括号匹配，如果能和栈顶匹配
	// 就将左边括号就行弹栈。 最后栈中数据为空。如果栈中数据不为空，说明左括号多了，验证失败。
	stack := []rune{}

	brackets := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}

	for _, char := range s {
		switch char {
		case '{', '(', '[':
			stack = append(stack, char)
		case ')', ']', '}':
			if len(stack) == 0 || stack[len(stack)-1] != brackets[char] {
				return false
			}
			// 弹栈
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
