package main

import "fmt"

/*
136. 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。
找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，
例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
*/
func singleNum(arr []int) int {
	var result int = -1e5
	if len(arr) == 0 {
		return result
	}
	// 2.声明一个计数map
	var countMap = make(map[int]int)

	// 3.遍历数组，记录每个元素出现次数
	for _, e := range arr {
		count := countMap[e]
		if count == 0 {
			countMap[e] = 1
		} else {
			countMap[e] = count + 1
		}
	}

	// 4.遍历计数map,找出计数为 1 的元素。
	flagCount := 1
	for k, v := range countMap {
		if v == flagCount {
			result = k
		}
	}
	return result
}

func main() {
	// 1.声明一个非空整数数组
	var arr = [7]int{-1, -1, 4, 4, 5, 6, 6}

	// 测试只出现一次的数字
	num := singleNum(arr[:])
	fmt.Println("====计数为【1】的元素====：", num)
}
